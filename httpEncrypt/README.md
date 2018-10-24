## httpEncry 使用fangfa

    1、go get -u -v git.dian.so/leto/util

    2、import(
        "git.dian.so/leto/util/httpEncrypt"
    )

    3、使用

        httpEncrypt.NewApp (source, secret, salt string) *app

        httpEncrypt.Get (ap *app, endpoint, url string, param interface{}, ver version, resp interface{}) error

        httpEncrypt.Post (ap *app, endpoint, url string, param interface{}, ver version, resp interface{}) error
