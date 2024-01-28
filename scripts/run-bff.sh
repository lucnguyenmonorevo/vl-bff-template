#!/bin/sh
DIR="$( cd "$( dirname "$0" )" && pwd )"

cd $DIR && cd ../ &&
./vl-bff-account/scripts/run.sh

cd $DIR && cd ../ &&
./vl-bff-business/scripts/run.sh

cd $DIR && cd ../ &&
./vl-bff-general-info/scripts/run.sh

cd $DIR && cd ../ &&
./vl-bff-notification/scripts/run.sh

cd $DIR && cd ../ &&
./vl-bff-physical-obj/scripts/run.sh

cd $DIR && cd ../ &&
./vl-bff-procurement/scripts/run.sh

cd $DIR && cd ../ &&
./vl-bff-production/scripts/run.sh
