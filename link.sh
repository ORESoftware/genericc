#!/usr/bin/env bash

dir="$(cd `dirname "$0"` && pwd)"


ln -sf "$dir/src/"* "$dir/lang/golang/src/genericc"


#ln -sf "$dir/src/greeter.c" "$dir/lang/golang/src/genericc/greeter.c"
#ln -sf "$dir/src/greeter.h" "$dir/lang/golang/src/genericc/greeter.h"

#ln -sf src/* lang/golang/src/main