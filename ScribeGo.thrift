include "fb303.thrift"

namespace go facebook.scribe

enum ResultCode
{
    OK,
    TRY_LATER
}

struct LogEntry
{
    1:  string category,
    2:  string message
}

service scribe extends fb303.FacebookService
{
    ResultCode Log(1: list<LogEntry> messages);
}
