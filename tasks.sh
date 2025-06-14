#!/bin/bash
set -e

# Usage: ./tasks.sh <command> -os <windows|linux> -env <local|dev|prod>

if [ -z "$1" ]; then
    echo "Usage: $0 <command> -os <windows|linux> -env <local|dev|prod>"
    exit 1
fi

CMD="$1"
shift

PLATFORM=""
ENV=""

# Parse flags
while [[ $# -gt 0 ]]; do
    case "$1" in
        -os)
            PLATFORM="$2"
            shift 2
            ;;
        -env)
            ENV="$2"
            shift 2
            ;;
        *)
            shift
            ;;
    esac
done

if [ -z "$PLATFORM" ]; then
    echo "Error: Platform not specified. Use -os windows or -os linux."
    exit 1
fi
if [ -z "$ENV" ]; then
    echo "Error: Environment not specified. Use -env local, dev, or prod."
    exit 1
fi

SCRIPT_DIR="commands/$PLATFORM"
SCRIPT_NAME="$CMD.$ENV"

# Try .sh, .bat, .cmd, or extensionless, in that order
SCRIPT=""
if [ -f "$SCRIPT_DIR/$SCRIPT_NAME.sh" ]; then
    SCRIPT="$SCRIPT_DIR/$SCRIPT_NAME.sh"
elif [ -f "$SCRIPT_DIR/$SCRIPT_NAME.bat" ]; then
    SCRIPT="$SCRIPT_DIR/$SCRIPT_NAME.bat"
elif [ -f "$SCRIPT_DIR/$SCRIPT_NAME.cmd" ]; then
    SCRIPT="$SCRIPT_DIR/$SCRIPT_NAME.cmd"
elif [ -f "$SCRIPT_DIR/$SCRIPT_NAME" ]; then
    SCRIPT="$SCRIPT_DIR/$SCRIPT_NAME"
fi

if [ -z "$SCRIPT" ]; then
    echo "Error: Command script not found: $SCRIPT_DIR/$SCRIPT_NAME.[sh|bat|cmd]"
    exit 1
fi

echo "[TRACE] Running: $SCRIPT"

# If .sh, run with bash; if .bat/.cmd, show how to run on Windows
case "$SCRIPT" in
    *.sh)
        bash "$SCRIPT"
        ;;
    *.bat|*.cmd)
        echo "To run on Windows: $SCRIPT"
        ;;
    *)
        chmod +x "$SCRIPT"
        "$SCRIPT"
        ;;
esac
