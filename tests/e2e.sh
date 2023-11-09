#!/usr/bin/env bash

set -ex

cp ginkgolinter testdata/src/a
cd testdata/src/a

# no suppress
[[ $(./ginkgolinter a/... 2>&1 | wc -l) == 2534 ]]
# suppress all but nil
[[ $(./ginkgolinter --suppress-len-assertion=true --suppress-err-assertion=true --suppress-compare-assertion=true --suppress-async-assertion=true --suppress-type-compare-assertion=true a/... 2>&1 | wc -l) == 1518 ]]
# suppress all but len
[[ $(./ginkgolinter --suppress-nil-assertion=true --suppress-err-assertion=true --suppress-compare-assertion=true --suppress-async-assertion=true --suppress-type-compare-assertion=true a/... 2>&1 | wc -l) == 816 ]]
# suppress all but err
[[ $(./ginkgolinter --suppress-nil-assertion=true --suppress-len-assertion=true --suppress-compare-assertion=true --suppress-async-assertion=true --suppress-type-compare-assertion=true a/... 2>&1 | wc -l) == 278 ]]
# suppress all but compare
[[ $(./ginkgolinter --suppress-nil-assertion=true --suppress-err-assertion=true --suppress-len-assertion=true --suppress-async-assertion=true --suppress-type-compare-assertion=true a/... 2>&1 | wc -l) == 288 ]]
# suppress all but async
[[ $(./ginkgolinter --suppress-nil-assertion=true --suppress-err-assertion=true --suppress-len-assertion=true --suppress-compare-assertion=true --suppress-type-compare-assertion=true a/... 2>&1 | wc -l) == 185 ]]
# suppress all but focus
[[ $(./ginkgolinter --suppress-nil-assertion=true --suppress-err-assertion=true --suppress-len-assertion=true --suppress-compare-assertion=true --suppress-async-assertion=true --forbid-focus-container=true --suppress-type-compare-assertion=true a/... 2>&1 | wc -l) == 209 ]]
# suppress all but compare different types
[[ $(./ginkgolinter --suppress-nil-assertion=true --suppress-err-assertion=true --suppress-len-assertion=true --suppress-compare-assertion=true --suppress-compare-assertion=true a/... 2>&1 | wc -l) == 281 ]]
# allow HaveLen(0)
[[ $(./ginkgolinter --allow-havelen-0=true a/... 2>&1 | wc -l) == 2521 ]]
# suppress all - should only return the few non-suppressble
[[ $(./ginkgolinter --suppress-nil-assertion=true --suppress-len-assertion=true --suppress-err-assertion=true --suppress-compare-assertion=true --suppress-async-assertion=true --forbid-focus-container=false --suppress-type-compare-assertion=true a/... 2>&1 | wc -l) == 154 ]]
