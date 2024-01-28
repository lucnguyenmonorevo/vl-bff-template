#!/bin/sh
DIR="$( cd "$( dirname "$0" )" && pwd )"

cd $DIR && cd ../ &&
./vl-bff-account/scripts/set_up_all.sh

cd $DIR && cd ../ &&
./vl-bff-business/scripts/set_up_all.sh

cd $DIR && cd ../ &&
./vl-bff-general-info/scripts/set_up_all.sh

cd $DIR && cd ../ &&
./vl-bff-notification/scripts/set_up_all.sh

cd $DIR && cd ../ &&
./vl-bff-physical-obj/scripts/set_up_all.sh

cd $DIR && cd ../ &&
./vl-bff-procurement/scripts/set_up_all.sh

cd $DIR && cd ../ &&
./vl-bff-production/scripts/set_up_all.sh
