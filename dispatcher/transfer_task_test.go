package dispatcher_test

import (
	"errors"
	"net"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/glestaris/ice-clique/api"
	"github.com/glestaris/ice-clique/dispatcher"
	"github.com/glestaris/ice-clique/dispatcher/fakes"
	"github.com/glestaris/ice-clique/scheduler"
	"github.com/glestaris/ice-clique/transfer"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TransferTask", func() {
	var (
		t              *dispatcher.TransferTask
		fakeServer     *fakes.FakeInterruptible
		fakeTransferrer *fakes.FakeTransferrer
		transferSpec   transfer.TransferSpec
		fakeRegistry   *fakes.FakeApiRegistry
		priority       int
		logger         *logrus.Logger
	)

	BeforeEach(func() {
		fakeServer = new(fakes.FakeInterruptible)
		fakeTransferrer = new(fakes.FakeTransferrer)
		transferSpec = transfer.TransferSpec{
			IP:   net.ParseIP("92.168.12.19"),
			Port: 1245,
			Size: 10 * 1024 * 1024,
		}
		fakeRegistry = new(fakes.FakeApiRegistry)
		priority = 10
		logger = &logrus.Logger{
			Out:       GinkgoWriter,
			Level:     logrus.DebugLevel,
			Formatter: new(logrus.TextFormatter),
		}

		t = &dispatcher.TransferTask{
			Server:          fakeServer,
			Transferrer:      fakeTransferrer,
			TransferSpec:    transferSpec,
			Registry:        fakeRegistry,
			DesiredPriority: priority,
			Logger:          logger,
		}
	})

	It("should return the provided priority", func() {
		Expect(t.Priority()).To(Equal(priority))
	})

	It("should run the transfer", func() {
		t.Run()
		Expect(fakeTransferrer.TransferCallCount()).To(Equal(1))
		Expect(fakeTransferrer.TransferArgsForCall(0)).To(Equal(transferSpec))
	})

	It("should pause the server", func() {
		t.Run()
	})

	It("should resume the server", func() {
		t.Run()
		Expect(fakeServer.ResumeCallCount()).To(Equal(1))
	})

	Context("when the task is failing for a while", func() {
		BeforeEach(func() {
			fakeTransferrer.TransferReturns(
				transfer.TransferResults{}, errors.New("banana"),
			)
		})

		It("should not change state", func() {
			for i := 0; i < 100; i++ {
				t.Run()
				Expect(t.State()).To(Equal(scheduler.TaskStateReady))
			}

			fakeTransferrer.TransferReturns(transfer.TransferResults{}, nil)
			t.Run()
			Expect(t.State()).To(Equal(scheduler.TaskStateDone))
		})

		It("should change transfer state back to pending", func() {
			for i := 0; i < 100; i++ {
				t.Run()
				Expect(t.TransferState()).To(Equal(api.TransferStatePending))
			}

			fakeTransferrer.TransferReturns(transfer.TransferResults{}, nil)
			t.Run()
			Expect(t.TransferState()).To(Equal(api.TransferStateCompleted))
		})
	})

	Context("when the task succeeds with results", func() {
		var transferResults transfer.TransferResults

		BeforeEach(func() {
			transferResults = transfer.TransferResults{
				Duration:  time.Millisecond * 100,
				Checksum:  uint32(12),
				BytesSent: uint32(10 * 1024 * 1024),
			}
			fakeTransferrer.TransferReturns(transferResults, nil)
		})

		It("should register the transfer results to the registry", func() {
			t.Run()

			Expect(fakeRegistry.RegisterResultsCallCount()).To(Equal(1))
			ip, res := fakeRegistry.RegisterResultsArgsForCall(0)
			Expect(ip).To(Equal(transferSpec.IP))
			Expect(res.IP).To(Equal(transferSpec.IP))
			Expect(res.BytesSent).To(Equal(transferResults.BytesSent))
			Expect(res.Checksum).To(Equal(transferResults.Checksum))
			Expect(res.Duration).To(Equal(transferResults.Duration))
			Expect(res.Time).To(BeTemporally("~", time.Now(), time.Second))
		})
	})

	Context("when the task takes time", func() {
		var transferrerChan chan bool

		BeforeEach(func() {
			transferrerChan = make(chan bool)

			fakeTransferrer.TransferStub = func(
				_ transfer.TransferSpec,
			) (transfer.TransferResults, error) {
				transferrerChan <- true

				<-transferrerChan

				return transfer.TransferResults{}, nil
			}
		})

		It("should change the transfer state to running", func() {
			Expect(t.TransferState()).To(Equal(api.TransferStatePending))

			runChan := make(chan struct{})
			go func() {
				t.Run()

				close(runChan)
			}()

			Eventually(transferrerChan).Should(Receive())
			Expect(t.TransferState()).To(Equal(api.TransferStateRunning))

			close(transferrerChan)
			Eventually(runChan).Should(BeClosed())

			Expect(t.TransferState()).To(Equal(api.TransferStateCompleted))
		})
	})
})
