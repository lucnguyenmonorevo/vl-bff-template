#!/bin/sh
DIR="$( cd "$( dirname "$0" )" && pwd )"

xterm -title "App 1" -e "cd $DIR && cd ../ &&
./vl-bff-account/scripts/auto_git_push.sh"

cd $DIR && cd ../ &&
./vl-bff-business/scripts/auto_git_push.sh

cd $DIR && cd ../ &&
./vl-bff-general-info/scripts/auto_git_push.sh

cd $DIR && cd ../ &&
./vl-bff-notification/scripts/auto_git_push.sh

cd $DIR && cd ../ &&
./vl-bff-physical-obj/scripts/auto_git_push.sh

cd $DIR && cd ../ &&
./vl-bff-procurement/scripts/auto_git_push.sh

cd $DIR && cd ../ &&
./vl-bff-production/scripts/auto_git_push.sh
