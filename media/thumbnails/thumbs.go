package thumbnails

// var ffmpegFound = false
// func init() {
// 	_, err := exec.LookPath("ffmpeg")
// 	if err == nil {
// 		ffmpegFound = true
// 	}
// }
//
// // OpenFfmpeg reads the first frame of a video using ffmpeg.
// // TODO: figure out why this fails for so many video inputs.
// func OpenFfmpeg(r io.Reader, w io.Writer, size int) error {
// 	// md := fmt.Sprintf("scale=%v:%v:force_original_aspect_ratio=decrease", size, size)
// 	// args := []string{"-i", "pipe:", "-frames:v", "1", "-vf", md, "-q:v", "10", "pipe:.jpg"}
// 	args := []string{"-i", "pipe:", "-f", "image2pipe", "-frames:v", "1", "pipe:.jpg"}
// 	cmd := exec.Command("ffmpeg", args...)
//
// 	cmd.Stdin = r
// 	cmd.Stdout = w
// 	cmd.Stderr = os.Stderr
// 	return cmd.Run()
// }
//
// // TODO: rewrite all this dogshit to encapsulate and generalize it
//
// var workerCount = runtime.GOMAXPROCS(0)
//
// type job struct {
// 	src  string
// 	dst  string
// 	file *os.File
// 	End  chan any
// }
//
// var (
// 	openCh = make(chan job, workerCount*2)
// 	workCh = make(chan job)
//
// 	wGroup   = workerGroup{}
// 	shutdown = make(chan struct{})
// )
//
// func Wait() {
// 	for {
// 		if wGroup.Count() == 0 {
// 			return
// 		}
// 		time.Sleep(100 * time.Millisecond)
// 	}
// }
//
// var maxSize int
//
// func Initialize(Max int) {
// 	maxSize = Max
// 	SetWorkerCount(workerCount)
// 	go opener(openCh)
// }
//
// func opener(ch chan job) {
// 	for {
// 		j := <-ch
// 		f, err := os.Open(j.src)
// 		if err != nil {
// 			clog.Error("error opening file", "src", j.src, "dst", j.dst, "error", err)
// 			f.Close()
// 			continue
// 		}
//
// 		j.file = f
// 		workCh <- j
// 	}
// }
//
// func worker(ch chan job) {
// 	wGroup.Add(1)
// 	clog.Debug("starting thumbnail worker", "count", wGroup.Count())
// 	defer wGroup.Done()
// 	for {
// 		select {
// 		case <-shutdown:
// 			return
// 		default:
// 			j := <-ch
//
// 			clog.Info("creating thumbnail file", "dst", j.dst)
// 			tar, err := os.Create(j.dst)
// 			if err != nil {
// 				clog.Error("error creating thumbnail file", "dst", j.dst, "error", err)
// 				tar.Close()
// 				j.file.Close()
// 				j.End <- true
// 				close(j.End)
// 				continue
// 			}
//
// 			m := media.Thumbnail{
// 				Quality: 50,
// 				MaxSize: maxSize,
// 				Kind:    media.GetKind(filepath.Ext(j.src)),
// 				Src:     j.file,
// 				Dst:     tar,
// 			}
// 			m.Create()
//
// 			tar.Close()
// 			j.file.Close()
// 			j.End <- true
// 			close(j.End)
// 		}
// 	}
// }
//
// func QueueThumb(src, dst string) (End <-chan any, err error) {
// 	if err = fsio.EnsureDir(filepath.Dir(dst)); err != nil {
// 		clog.Error("error ensuring dst", "dst", filepath.Dir(dst), "error", err)
// 		return
// 	}
//
// 	if fsio.Exists(dst) {
// 		err = fmt.Errorf("thumbnail already exists")
// 		clog.Debug("thumbnail already exists", "src", src, "dst", dst)
// 		return
// 	}
//
// 	j := job{src: src, dst: dst, file: nil, End: make(chan any, 1)}
// 	openCh <- j
// 	return j.End, nil
// }
//
// type workerGroup struct{ count int64 }
//
// func (wg *workerGroup) Add(delta int) { atomic.AddInt64(&wg.count, int64(delta)) }
// func (wg *workerGroup) Done()         { atomic.AddInt64(&wg.count, -1) }
// func (wg *workerGroup) Count() int    { return int(atomic.LoadInt64(&wg.count)) }
//
// func SetWorkerCount(count int) {
// 	wc := wGroup.Count()
// 	if count <= 0 {
// 		if wGroup.Count() == 0 {
// 			return
// 		}
//
// 		for i := 0; i < wc; i++ {
// 			shutdown <- struct{}{}
// 		}
// 	}
//
// 	if count > wc {
// 		for i := 0; i < count-wc; i++ {
// 			go worker(workCh)
// 		}
// 	}
// }
