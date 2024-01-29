#!/bin/sh
DIR="$( cd "$( dirname "$0" )" && pwd )"
#!/bin/bash
killbg() {
        for p in "${pids[@]}" ; do
                kill "$p";
        done
}
trap killbg EXIT
pids=()

cd $DIR && cd ../ &&
./vl-bff-account/scripts/run.sh &
pids+=($!)
cd $DIR && cd ../ &&
./vl-bff-business/scripts/run.sh &
pids+=($!)
cd $DIR && cd ../ &&
./vl-bff-general-info/scripts/run.sh &
pids+=($!)
cd $DIR && cd ../ &&
./vl-bff-notification/scripts/run.sh &
pids+=($!)
cd $DIR && cd ../ &&
./vl-bff-physical-obj/scripts/run.sh &
pids+=($!)
cd $DIR && cd ../ &&
./vl-bff-procurement/scripts/run.sh &
pids+=($!)
cd $DIR && cd ../ &&
./vl-bff-production/scripts/run.sh &
pids+=($!)

sleep 5
cd $DIR && cd ../ &&
./vl-bff-router/scripts/run.sh