param([int] $p)

function time($block) {
    $sw = [Diagnostics.Stopwatch]::StartNew()
    &$block
    $sw.Stop()
    $sw.Elapsed.TotalSeconds
}

go build
time { .\euler.exe -p $p }
go clean
go fmt