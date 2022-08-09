package errof

type UserErr string
type InternalErr string

func (e UserErr) Error() (msg string) {
	var ok bool
	if msg, ok = ErrCodeNames[e]; !ok {
		return string(e)
	}
	return msg
}

func (e InternalErr) Error() (msg string) {
	var ok bool
	if msg, ok = InternalErrCodeNames[e]; !ok {
		return string(e)
	}
	return msg
}

// InternalErrCodeNames :
var InternalErrCodeNames = map[InternalErr]string{
	ErrInternal:          "内部エラーが発生しました。",
	ErrDatabase:          "データベースでの不整合が発生しました。",
	ErrDataInconsistency: "データの不整合が発生しました。",
	ErrWriteResponse:     "レスポンスの書き込みに失敗しました。",
	ErrReadResponse:      "レスポンスの読み込みに失敗しました。",
	ErrParseURL:          "URLの解析に失敗しました。",
	ErrNewHTTPClient:     "HTTPクライアントの生成に失敗しました。",
	ErrHTTPRequest:       "HTTPリクエストの実行に失敗しました。",
	ErrStatusCode:        "ステータスコードが正常ではありません。",
	ErrUnmarshalResponse: "レスポンスのアンマーシャリングに失敗しました。",
}

// CmdErrCodeNames :
var ErrCodeNames = map[UserErr]string{
	ErrNoPost: "post not found.",
}

var (
	// Internal Err
	ErrInternal          InternalErr = "ErrInternal"
	ErrDatabase          InternalErr = "ErrDatabase"
	ErrDataInconsistency InternalErr = "ErrDataInconsistency"
	ErrWriteResponse     InternalErr = "ErrWriteResponse"
	ErrReadResponse      InternalErr = "ErrReadResponse"
	ErrParseURL          InternalErr = "ErrParseURL"
	ErrNewHTTPClient     InternalErr = "ErrCreateHTTPRequestClient"
	ErrHTTPRequest       InternalErr = "ErrHTTPRequest"
	ErrStatusCode        InternalErr = "ErrStatusCode"
	ErrUnmarshalResponse InternalErr = "ErrUnmarshalResponse"

	// User Err
	ErrNoPost UserErr = "ErrNoPost"
)
