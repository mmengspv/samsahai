package exporter

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	"github.com/agoda-com/samsahai/internal"
	"github.com/agoda-com/samsahai/internal/util/http"
	"github.com/agoda-com/samsahai/internal/util/unittest"
	s2hv1beta1 "github.com/agoda-com/samsahai/pkg/apis/env/v1beta1"
)

func TestExporter(t *testing.T) {
	unittest.InitGinkgo(t, "Samsahai Exporter")
}

var cfg *rest.Config
var c client.Client

func TestMain(m *testing.M) {
	var err error
	t := &envtest.Environment{
		CRDDirectoryPaths: []string{filepath.Join("..", "..", "..", "config", "crds")},
	}

	err = s2hv1beta1.SchemeBuilder.AddToScheme(scheme.Scheme)
	if err != nil {
		log.Fatal(err)
	}

	if cfg, err = t.Start(); err != nil {
		logger.Error(err, "start testenv error")
		os.Exit(1)
	}

	if c, err = client.New(cfg, client.Options{Scheme: scheme.Scheme}); err != nil {
		log.Fatal(err)
	}

	code := m.Run()
	_ = t.Stop()
	os.Exit(code)
}

func startDate(y, mo, d, h, mi, s int) *metav1.Time {
	ti := metav1.Date(y, time.Month(mo), d, h, mi, s, 0, time.UTC)
	return &ti
}
func date(y, mo, d, h, mi, s int) metav1.Time {
	ti := metav1.Date(y, time.Month(mo), d, h, mi, s, 0, time.UTC)
	return ti
}

var _ = Describe("Samsahai Exporter", func() {
	timeout := float64(3000)
	namespace := "default"
	g := NewWithT(GinkgoT())
	var wgStop *sync.WaitGroup
	var chStop chan struct{}
	var SamsahaiURL = "aaa"

	RegisterMetrics()

	BeforeEach(func(done Done) {
		defer close(done)
		defer GinkgoRecover()

		chStop = make(chan struct{})

		mgr, err := manager.New(cfg, manager.Options{Namespace: namespace, MetricsBindAddress: ":8008"})
		Expect(err).NotTo(HaveOccurred(), "should create manager successfully")

		t := map[string]internal.ConfigManager{
			"example":      nil,
			"testTeamname": nil,
		}
		q := &s2hv1beta1.QueueList{
			Items: []s2hv1beta1.Queue{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "qName1",
						Namespace: namespace,
					},
					Spec: s2hv1beta1.QueueSpec{
						TeamName: "testQTeamName1",
						Version:  "10.9.8.7",
					},
					Status: s2hv1beta1.QueueStatus{
						NoOfProcessed: 1,
						State:         "testQState1",
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "qName2",
						Namespace: namespace,
					},
					Spec: s2hv1beta1.QueueSpec{
						TeamName: "testQTeamName2",
						Version:  "7.8.9.10",
					},
					Status: s2hv1beta1.QueueStatus{
						NoOfProcessed: 1,
						State:         "testQState2",
					},
				},
			},
		}

		qh := &s2hv1beta1.QueueHistoryList{
			Items: []s2hv1beta1.QueueHistory{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "testQHname1",
						Namespace: namespace,
					},
					Spec: s2hv1beta1.QueueHistorySpec{
						Queue: &s2hv1beta1.Queue{
							Spec: s2hv1beta1.QueueSpec{
								TeamName: "testQHTeamName1",
								Version:  "1.2.3.4",
							},
						},
						IsDeploySuccess: true,
						IsTestSuccess:   true,
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "testQHname2",
						Namespace: namespace,
					},
					Spec: s2hv1beta1.QueueHistorySpec{
						Queue: &s2hv1beta1.Queue{
							Spec: s2hv1beta1.QueueSpec{
								TeamName: "testQHTeamName2",
								Version:  "4.3.2.1",
							},
						},
						IsDeploySuccess: true,
						IsTestSuccess:   true,
					},
				},
			},
		}

		ap := &s2hv1beta1.ActivePromotionList{
			Items: []s2hv1beta1.ActivePromotion{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "testAPName1",
						Namespace: namespace,
					},
					Status: s2hv1beta1.ActivePromotionStatus{
						State: s2hv1beta1.ActivePromotionWaiting,
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "testAPName2",
						Namespace: namespace,
					},
					Status: s2hv1beta1.ActivePromotionStatus{
						State: s2hv1beta1.ActivePromotionFinished,
					},
				},
			},
		}

		aph := &s2hv1beta1.ActivePromotionHistoryList{
			Items: []s2hv1beta1.ActivePromotionHistory{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "testAPHName1",
						Namespace: namespace,
						Labels: map[string]string{
							"samsahai.io/teamname": "testAPHTeamname1",
						},
					},
					Spec: s2hv1beta1.ActivePromotionHistorySpec{
						ActivePromotion: &s2hv1beta1.ActivePromotion{
							Status: s2hv1beta1.ActivePromotionStatus{
								State:     "DestroyingPreviousActiveEnvironment",
								StartedAt: startDate(2019, 12, 10, 2, 22, 02),
								PreActiveQueue: s2hv1beta1.QueueStatus{
									Conditions: []s2hv1beta1.QueueCondition{
										{
											Type:               "QueueDeployed",
											LastTransitionTime: date(2019, 12, 10, 3, 3, 36),
										},
									},
								},
								Conditions: []s2hv1beta1.ActivePromotionCondition{
									{
										Type:               "ActivePromotionStarted",
										LastTransitionTime: date(2019, 12, 10, 2, 22, 02),
									},
									{
										Type:               "PreActiveVerified",
										LastTransitionTime: date(2019, 12, 10, 3, 38, 21),
									},
									{
										Type:               "ActivePromoted",
										LastTransitionTime: date(2019, 12, 10, 3, 39, 01),
									},
									{
										Type:               "Finished",
										LastTransitionTime: date(2019, 12, 10, 4, 11, 13),
									},
								},
								Result: s2hv1beta1.ActivePromotionSuccess,
							},
						},
					},
				},
			},
		}
		oc := &s2hv1beta1.ActivePromotion{
			ObjectMeta: metav1.ObjectMeta{
				Name: "testOCTeamName",
			},
			Status: s2hv1beta1.ActivePromotionStatus{
				OutdatedComponents: []*s2hv1beta1.OutdatedComponent{
					{
						Name: "testOCName1",
						CurrentImage: &s2hv1beta1.Image{
							Tag: "2019.12.07.00",
						},
						LatestImage: &s2hv1beta1.Image{
							Tag: "2019.12.10.00",
						},
						OutdatedDuration: 99540000000000,
					},
					{
						Name: "testOCName2",
						CurrentImage: &s2hv1beta1.Image{
							Tag: "2019.12.07.00",
						},
						LatestImage: &s2hv1beta1.Image{
							Tag: "2019.12.10.00",
						},
						OutdatedDuration: 99599999999999,
					},
				},
			},
		}

		SetTeamNameMetric(t)
		SetQueueMetric(q)
		SetQueueHistoriesMetric(qh, SamsahaiURL)
		SetActivePromotionMetric(ap)
		SetActivePromotionHistoriesMetric(aph)
		SetOutdatedComponentMetric(oc)
		SetHealthStatusMetric("9.9.9.8", "777888999", 234000)

		wgStop = &sync.WaitGroup{}
		wgStop.Add(1)
		go func() {
			defer wgStop.Done()
			Expect(mgr.Start(chStop)).To(BeNil())
		}()
	}, timeout)

	AfterEach(func(done Done) {
		defer close(done)
		close(chStop)
		wgStop.Wait()
	}, timeout)

	It("Should show team name correctly ", func() {
		data, err := http.Get("http://localhost:8008/metrics")
		g.Expect(err).NotTo(HaveOccurred())
		expectedData := strings.Contains(string(data), `samsahai_team{teamName="example"} 1`)
		g.Expect(expectedData).To(BeTrue())
	}, timeout)

	It("Should show queue metric correctly  ", func(done Done) {
		defer close(done)
		data, err := http.Get("http://localhost:8008/metrics")
		g.Expect(err).NotTo(HaveOccurred())
		expectedData := strings.Contains(string(data), `samsahai_queue{component="qName1",no_of_processed="1",order="0",state="testQState1",teamName="testQTeamName1",version="10.9.8.7"} 1`)
		g.Expect(expectedData).To(BeTrue())
		expectedData = strings.Contains(string(data), `samsahai_queue{component="qName2",no_of_processed="1",order="0",state="testQState2",teamName="testQTeamName2",version="7.8.9.10"} 1`)
		g.Expect(expectedData).To(BeTrue())
		expectedData = strings.Contains(string(data), `samsahai_queue{component="",`)
		g.Expect(expectedData).To(BeFalse())
	}, timeout)

	It("Should show queue histories metric correctly ", func(done Done) {
		defer close(done)
		data, err := http.Get("http://localhost:8008/metrics")
		g.Expect(err).NotTo(HaveOccurred())
		expectedData := strings.Contains(string(data), `samsahai_queue_histories{component="testQHname1",log="aaa/team/testQHTeamName1/queue/histories/testQHname1/log",result="success",teamName="testQHTeamName1",version="1.2.3.4"} 1`)
		g.Expect(expectedData).To(BeTrue())
		expectedData = strings.Contains(string(data), `samsahai_queue_histories{component="testQHname2",log="aaa/team/testQHTeamName2/queue/histories/testQHname2/log",result="success",teamName="testQHTeamName2",version="4.3.2.1"} 1`)
		g.Expect(expectedData).To(BeTrue())
		expectedData = strings.Contains(string(data), `samsahai_queue_histories{component="",`)
		g.Expect(expectedData).To(BeFalse())
	}, timeout)

	It("Should show active promotion correctly", func(done Done) {
		defer close(done)
		data, err := http.Get("http://localhost:8008/metrics")
		g.Expect(err).NotTo(HaveOccurred())
		expectedData := strings.Contains(string(data), `samsahai_active_promotion{status="Finished",teamName="testAPName2"} 1`)
		g.Expect(expectedData).To(BeTrue())
		expectedData = strings.Contains(string(data), `samsahai_active_promotion{status="Waiting",teamName="testAPName1"} 1`)
		g.Expect(expectedData).To(BeTrue())

	}, timeout)

	It("Should show active promotion histories correctly", func(done Done) {
		defer close(done)
		data, err := http.Get("http://localhost:8008/metrics")
		g.Expect(err).NotTo(HaveOccurred())
		expectedData := strings.Contains(string(data), `samsahai_active_promotion_histories{name="testAPHName1",result="Success",startTime="2019-12-10T02:22:02Z",state="deploying",teamName="testAPHTeamname1"} 2494`)
		g.Expect(expectedData).To(BeTrue())
		expectedData = strings.Contains(string(data), `samsahai_active_promotion_histories{name="testAPHName1",result="Success",startTime="2019-12-10T02:22:02Z",state="destroying",teamName="testAPHTeamname1"} 1932`)
		g.Expect(expectedData).To(BeTrue())
		expectedData = strings.Contains(string(data), `samsahai_active_promotion_histories{name="testAPHName1",result="Success",startTime="2019-12-10T02:22:02Z",state="promoting",teamName="testAPHTeamname1"} 40`)
		g.Expect(expectedData).To(BeTrue())
		expectedData = strings.Contains(string(data), `samsahai_active_promotion_histories{name="testAPHName1",result="Success",startTime="2019-12-10T02:22:02Z",state="testing",teamName="testAPHTeamname1"} 2085`)
		g.Expect(expectedData).To(BeTrue())
		expectedData = strings.Contains(string(data), `samsahai_active_promotion_histories{name="testAPHName1",result="Success",startTime="2019-12-10T02:22:02Z",state="waiting",teamName="testAPHTeamname1"} 0`)
		g.Expect(expectedData).To(BeTrue())
	}, timeout)

	It("Should show outdated component correctly", func(done Done) {
		defer close(done)
		data, err := http.Get("http://localhost:8008/metrics")
		g.Expect(err).NotTo(HaveOccurred())
		expectedData := strings.Contains(string(data), `samsahai_outdated_component{component="testOCName1",currentVer="2019.12.07.00",desiredVer="2019.12.10.00",teamName="testOCTeamName"} 2.432554678116942e+18`)
		g.Expect(expectedData).To(BeTrue())
		expectedData = strings.Contains(string(data), `samsahai_outdated_component{component="testOCName2",currentVer="2019.12.07.00",desiredVer="2019.12.10.00",teamName="testOCTeamName"} 4.214431078557532e+17`)
		g.Expect(expectedData).To(BeTrue())
	}, timeout)

	It("Should show health metric correctly", func(done Done) {
		defer close(done)
		data, err := http.Get("http://localhost:8008/metrics")
		g.Expect(err).NotTo(HaveOccurred())
		expectedData := strings.Contains(string(data), `samsahai_health{gitCommit="777888999",version="9.9.9.8"} 234000`)
		g.Expect(expectedData).To(BeTrue())
	}, timeout)

})
