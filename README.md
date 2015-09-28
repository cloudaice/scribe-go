Over View
=========

scribe 的Go语言客户端。

Detail
======

在scribe的官方项目中只提供了Python版本的客户端，并且在扩展说明里面提到了Java版本的客户端。
这里根据scribe的thrift定义，使用0.9.1版本thrift进行编译生成`fb303`和`scribe`两个库。
并且在这两个库的基础上封装了一个scribe的Go语言客户端，实现向scribe server 发送消息的功能。

Getting Started
===============

    import "github.com/cloudaice/scribe-go/scribe"
    

    client, err := scribe.NewScribeLoger("host", "port")
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()

    ok, err := client.SendOne("category", "message")
    if err != nil {
        fmt.Println(err)
    } else if ok {
        fmt.Println("ok")
    } else {
        fmt.Println("try later")
    }

也可以直接使用编译出来的scribe包，自主定义

    import "github.com/cloudaice/scribe-go/facebook/scribe"
    
