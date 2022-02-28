package gupshup

import (
	"bytes"
	"errors"
	"github.com/jhidalgoesp/gupshup-whatsapp-go/mocks"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

var (
	textRequest = TextRequest{
		Source:      "123456789",
		Destination: "123456789",
		Text:        "lorem",
	}
	imageRequest = ImageRequest{
		Source:      "123456789",
		Destination: "123456789",
		Image:       "https://www.buildquickbots.com/whatsapp/media/sample/jpg/sample01.jpg",
		Preview:     "https://www.buildquickbots.com/whatsapp/media/sample/jpg/sample01.jpg",
		Caption:     "Caption",
	}
	documentRequest = DocumentRequest{
		Source:      "123456789",
		Destination: "123456789",
		Url:         "https://www.buildquickbots.com/whatsapp/media/sample/pdf/sample01.pdf",
		Filename:    "Sample funtional resume",
	}
	audioRequest = AudioRequest{
		Url: "https://www.buildquickbots.com/whatsapp/media/sample/audio/sample01.mp3",
	}
	videoRequest = VideoRequest{
		Url:     "https://www.buildquickbots.com/whatsapp/media/sample/video/sample01.mp4",
		Caption: "Sample video",
	}
	stickerRequest = StickerRequest{
		Source:      "123456789",
		Destination: "123456789",
		Url:         "http://www.buildquickbots.com/whatsapp/stickers/SampleSticker01.webp",
	}
	interactiveMessageRequest = InteractiveMessageRequest{
		Source:      "593958711086",
		Destination: "593988434471",
		InteractiveMessage: NewInteractiveMessage(
			"Body",
			"Title",
			"123456",
			[]InteractiveGlobalButtons{
				NewButton("Button"),
			},
			[]InteractiveMessageItem{
				NewInteractiveMessageItem(
					"first section",
					"first section subtitle",
					[]InteractiveMessageOptions{
						NewInteractiveMessageOption("1st option", "1st option description", "1"),
						NewInteractiveMessageOption("2nd option", "2nd option description", "2"),
						NewInteractiveMessageOption("3rd option", "3rd option description", "3"),
					},
				),
				NewInteractiveMessageItem(
					"second section",
					"second section subtitle",
					[]InteractiveMessageOptions{
						NewInteractiveMessageOption("4st option", "4st option description", "4"),
						NewInteractiveMessageOption("5nd option", "5nd option description", "5"),
						NewInteractiveMessageOption("6rd option", "6rd option description", "6"),
					},
				),
			},
		),
	}
	gc = Client{ApiKey: "test", AppName: "test"}
)

func init() {
	HttpClient = &mocks.MockClient{}
	Http = &mocks.MockHttp{}
}

func mockGetDoFunc() {
	response := `{"status":"submitted","messageId":"1234-56789"}`
	responseBody := ioutil.NopCloser(bytes.NewReader([]byte(response)))
	mocks.GetDoFunc = func(r *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       responseBody,
		}, nil
	}
}

func mockGetBuildRequestFunc() {
	mocks.GetBuildRequestFunc = func(method, url string, body io.Reader) (*http.Request, error) {
		return http.NewRequest(method, url, body)
	}
}

func assertResponse(t testing.TB, err error, got, want Response) {
	t.Helper()
	if err != nil {
		t.Errorf("didn't expected an error")
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v wanted %v", got, want)
	}
}

func assertError(t testing.TB, err error, got, want Response) {
	t.Helper()
	if err == nil {
		t.Errorf("expected an error")
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("text message sent successfully")
	}
}

func TestGupshupClient_SendText(t *testing.T) {
	t.Run("text message sent successfully", func(t *testing.T) {
		mockGetDoFunc()
		mockGetBuildRequestFunc()
		got, err := gc.SendText(textRequest)
		want := Response{"submitted", "1234-56789"}
		assertResponse(t, err, got, want)
	})

	t.Run("error when using wrong gupshup url", func(t *testing.T) {
		mocks.GetBuildRequestFunc = func(method, url string, body io.Reader) (*http.Request, error) {
			return nil, errors.New("error wrong url")
		}
		got, err := gc.SendText(textRequest)
		want := Response{}
		assertError(t, err, got, want)
	})

	t.Run("error when sending text message", func(t *testing.T) {
		mockGetBuildRequestFunc()
		mocks.GetDoFunc = func(r *http.Request) (*http.Response, error) {
			return nil, errors.New("error from gupshup")
		}
		got, err := gc.SendText(textRequest)
		want := Response{}
		assertError(t, err, got, want)
	})
}

func TestGupshupClient_SendImage(t *testing.T) {
	t.Run("image message sent successfully", func(t *testing.T) {
		mockGetDoFunc()
		mockGetBuildRequestFunc()
		got, err := gc.SendImage(imageRequest)
		want := Response{"submitted", "1234-56789"}
		assertResponse(t, err, got, want)
	})

	t.Run("error when using wrong gupshup url", func(t *testing.T) {
		mocks.GetBuildRequestFunc = func(method, url string, body io.Reader) (*http.Request, error) {
			return nil, errors.New("error wrong url")
		}
		got, err := gc.SendImage(imageRequest)
		want := Response{}
		assertError(t, err, got, want)
	})

	t.Run("error when sending image message", func(t *testing.T) {
		mockGetBuildRequestFunc()
		mocks.GetDoFunc = func(r *http.Request) (*http.Response, error) {
			return nil, errors.New("error from gupshup")
		}
		got, err := gc.SendImage(imageRequest)
		want := Response{}
		assertError(t, err, got, want)
	})
}

func TestGupshupClient_SendDocument(t *testing.T) {
	t.Run("document message sent successfully", func(t *testing.T) {
		mockGetDoFunc()
		mockGetBuildRequestFunc()
		got, err := gc.SendDocument(documentRequest)
		want := Response{"submitted", "1234-56789"}
		assertResponse(t, err, got, want)
	})

	t.Run("error when using wrong gupshup url", func(t *testing.T) {
		mocks.GetBuildRequestFunc = func(method, url string, body io.Reader) (*http.Request, error) {
			return nil, errors.New("error wrong url")
		}
		got, err := gc.SendDocument(documentRequest)
		want := Response{}
		assertError(t, err, got, want)
	})

	t.Run("error when sending document message", func(t *testing.T) {
		mockGetBuildRequestFunc()
		mocks.GetDoFunc = func(r *http.Request) (*http.Response, error) {
			return nil, errors.New("error from gupshup")
		}
		got, err := gc.SendDocument(documentRequest)
		want := Response{}
		assertError(t, err, got, want)
	})
}

func TestGupshupClient_SendAudio(t *testing.T) {
	t.Run("audio message sent successfully", func(t *testing.T) {
		mockGetDoFunc()
		mockGetBuildRequestFunc()
		got, err := gc.SendAudio(audioRequest)
		want := Response{"submitted", "1234-56789"}
		assertResponse(t, err, got, want)
	})

	t.Run("error when using wrong gupshup url", func(t *testing.T) {
		mocks.GetBuildRequestFunc = func(method, url string, body io.Reader) (*http.Request, error) {
			return nil, errors.New("error wrong url")
		}
		got, err := gc.SendAudio(audioRequest)
		want := Response{}
		assertError(t, err, got, want)
	})

	t.Run("error when sending document message", func(t *testing.T) {
		mockGetBuildRequestFunc()
		mocks.GetDoFunc = func(r *http.Request) (*http.Response, error) {
			return nil, errors.New("error from gupshup")
		}
		got, err := gc.SendAudio(audioRequest)
		want := Response{}
		assertError(t, err, got, want)
	})
}

func TestGupshupClient_SendVideo(t *testing.T) {
	t.Run("video message sent successfully", func(t *testing.T) {
		mockGetDoFunc()
		mockGetBuildRequestFunc()
		got, err := gc.SendVideo(videoRequest)
		want := Response{"submitted", "1234-56789"}
		assertResponse(t, err, got, want)
	})

	t.Run("error when using wrong gupshup url", func(t *testing.T) {
		mocks.GetBuildRequestFunc = func(method, url string, body io.Reader) (*http.Request, error) {
			return nil, errors.New("error wrong url")
		}
		got, err := gc.SendVideo(videoRequest)
		want := Response{}
		assertError(t, err, got, want)
	})

	t.Run("error when sending document message", func(t *testing.T) {
		mockGetBuildRequestFunc()
		mocks.GetDoFunc = func(r *http.Request) (*http.Response, error) {
			return nil, errors.New("error from gupshup")
		}
		got, err := gc.SendVideo(videoRequest)
		want := Response{}
		assertError(t, err, got, want)
	})
}

func TestGupshupClient_SendSticker(t *testing.T) {
	t.Run("sticker message sent successfully", func(t *testing.T) {
		mockGetDoFunc()
		mockGetBuildRequestFunc()
		got, err := gc.SendSticker(stickerRequest)
		want := Response{"submitted", "1234-56789"}
		assertResponse(t, err, got, want)
	})

	t.Run("error when using wrong gupshup url", func(t *testing.T) {
		mocks.GetBuildRequestFunc = func(method, url string, body io.Reader) (*http.Request, error) {
			return nil, errors.New("error wrong url")
		}
		got, err := gc.SendSticker(stickerRequest)
		want := Response{}
		assertError(t, err, got, want)
	})

	t.Run("error when sending document message", func(t *testing.T) {
		mockGetBuildRequestFunc()
		mocks.GetDoFunc = func(r *http.Request) (*http.Response, error) {
			return nil, errors.New("error from gupshup")
		}
		got, err := gc.SendSticker(stickerRequest)
		want := Response{}
		assertError(t, err, got, want)
	})
}

func TestGupshupClient_SendInteractiveMessage(t *testing.T) {
	t.Run("interactive message sent successfully", func(t *testing.T) {
		mockGetDoFunc()
		mockGetBuildRequestFunc()
		got, err := gc.SendInteractiveMessage(interactiveMessageRequest)
		want := Response{"submitted", "1234-56789"}
		assertResponse(t, err, got, want)
	})

	t.Run("error when using wrong gupshup url", func(t *testing.T) {
		mocks.GetBuildRequestFunc = func(method, url string, body io.Reader) (*http.Request, error) {
			return nil, errors.New("error wrong url")
		}
		got, err := gc.SendInteractiveMessage(interactiveMessageRequest)
		want := Response{}
		assertError(t, err, got, want)
	})

	t.Run("error when sending document message", func(t *testing.T) {
		mockGetBuildRequestFunc()
		mocks.GetDoFunc = func(r *http.Request) (*http.Response, error) {
			return nil, errors.New("error from gupshup")
		}
		got, err := gc.SendInteractiveMessage(interactiveMessageRequest)
		want := Response{}
		assertError(t, err, got, want)
	})
}
