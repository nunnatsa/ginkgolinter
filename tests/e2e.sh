#!/usr/bin/env bash

set -ex

cp ginkgolinter testdata/src/a
cd testdata/src/a

# no suppress
[[ $(./ginkgolinter a/... 2>&1 | wc -l) == 2296 ]]
# suppress all but nil
[[ $(./ginkgolinter --suppress-len-assertion=true --suppress-err-assertion=true --suppress-compare-assertion=true --suppress-async-assertion=true a/... 2>&1 | wc -l) == 1390 ]]
# suppress all but len
[[ $(./ginkgolinter --suppress-nil-assertion=true --suppress-err-assertion=true --suppress-compare-assertion=true --suppress-async-assertion=true a/... 2>&1 | wc -l) == 675 ]]
# suppress all but err
[[ $(./ginkgolinter --suppress-nil-assertion=true --suppress-len-assertion=true --suppress-compare-assertion=true --suppress-async-assertion=true a/... 2>&1 | wc -l) == 151 ]]
# suppress all but compare
[[ $(./ginkgolinter --suppress-nil-assertion=true --suppress-err-assertion=true --suppress-len-assertion=true --suppress-async-assertion=true a/... 2>&1 | wc -l) == 148 ]]
# suppress all but async
[[ $(./ginkgolinter --suppress-nil-assertion=true --suppress-err-assertion=true --suppress-len-assertion=true --suppress-compare-assertion=true a/... 2>&1 | wc -l) == 58 ]]
# allow HaveLen(0)
[[ $(./ginkgolinter --allow-havelen-0=true a/... 2>&1 | wc -l) == 2283 ]]
# suppress all - should do nothing
[[ $(./ginkgolinter --suppress-nil-assertion=true --suppress-len-assertion=true --suppress-err-assertion=true --suppress-compare-assertion=true --suppress-async-assertion=true a/... 2>&1 | wc -l) == 0 ]]
