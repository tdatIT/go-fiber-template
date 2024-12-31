# GPLog

**GPLog** is a library used to print logs in a specified format that supports centralized logging. It is developed based on the `uber-zap` library.

![Example Import](example.png)

### Log Format

The log output follows this format:

```json
{"level":"warn","time":"2024-09-11T17:37:18.378+0700","caller":"v2@v2.52.5/router.go:145","func":"github.com/gofiber/fiber/v2.(*App).next","msg":"Client error","service_name":"example-service","ip":"127.0.0.1","latency":"19.334µs","status":404,"method":"GET","url":"/sdsad"}
```

```json
{"level":"info","time":"2024-09-11T17:36:52.743+0700","caller":"example/main.go:16","func":"main.main","msg":"Hello","service_name":"example-service"}
```
### Installation
To install, use the following command:

```bash
make install
```
### Usage
You can either create a new instance of GPLog or use the singleton instance that is initialized when the program starts.

Creating a New Logger Instance
```go
gpLog := NewLogger(&LogConfig{
    Level:      "debug",
    LogFormat:  JsonFormat,
    TimeFormat: ISO8601TimeEncoder,
    Filename:   "",
    ServiceName: "example-service",
})
```
You can also use the default configuration:

```go
gpLog := NewLogger(DefaultConfig)
```
The singleton logger supports both SugaredLogger for formatted output and Zap's regular logger for structured logging.

```go

gplog.Debugf("Hello %s", "world")
gplog.Debug("Hello world", zap.String("key", "value"))
gplog.Infof("Hello %s", "world")
gplog.Info("Hello world", zap.String("key", "value"))
gplog.Warnf("Hello %s", "world")
gplog.Warn("Hello world", zap.String("key", "value"))
gplog.Errorf("Hello %s", "world")
gplog.Error("Hello world", zap.String("key", "value"))
gplog.DPanicf("Hello %s", "world")
gplog.DPanic("Hello world", zap.String("key", "value"))
gplog.Panicf("Hello %s", "world")
gplog.Panic("Hello world", zap.String("key", "value"))
gplog.Fatalf("Hello %s", "world")
gplog.Fatal("Hello world", zap.String("key", "value"))
```

Updating the Singleton Logger

```go
newgplog := gplog.NewLogger(&gplog.LogConfig{
    ServiceName: "example-service",
    Level:       "Error",
    LogFormat:   gplog.ConsoleFormat,
    TimeFormat:  gplog.RFC3339NanoTimeEncoder,
    Filename:    "trace.log",
})
gplog.SetLogger(newgplog)
```

### Integrating with Frameworks
For frameworks that can integrate with uber-zap, you can use the initialized logger. Below is an example using Fiber-v2:

```go
newgplog := gplog.NewLogger(&gplog.LogConfig{
    ServiceName: "example-service",
    Level:       "Error",
    LogFormat:   gplog.ConsoleFormat,
    TimeFormat:  gplog.RFC3339NanoTimeEncoder,
    Filename:    "trace.log",
})

app := fiber.New()
app.Use(fiberzap.New(fiberzap.Config{
    Logger: newgplog.GetZapInstance(),
}))

app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Hello, World!")
})

app.Listen(":3000")

```
For the singleton instance, you can access the Zap instance with ```gplog.GetZapInstance()```

<hr>
Developed by tdat.it2k2@gmail.com