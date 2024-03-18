#!/usr/bin/env bash

set -ex

cp ginkgolinter testdata/src/a
cd testdata/src/a

# no suppress
[[ $(./ginkgolinter a/... 2>&1 | wc -l) == 2591 ]]
# suppress all but nil
[[ $(./ginkgolinter --suppress-len-assertion=true --suppress-err-assertion=true --suppress-compare-assertion=true --suppress-async-assertion=true --suppress-type-compare-assertion=true a/... 2>&1 | wc -l) == 1516 ]]
# suppress all but len
[[ $(./ginkgolinter --suppress-nil-assertion=true --suppress-err-assertion=true --suppress-compare-assertion=true --suppress-async-assertion=true --suppress-type-compare-assertion=true a/... 2>&1 | wc -l) == 876 ]]
# suppress all but err
[[ $(./ginkgolinter --suppress-nil-assertion=true --suppress-len-assertion=true --suppress-compare-assertion=true --suppress-async-assertion=true --suppress-type-compare-assertion=true a/... 2>&1 | wc -l) == 276 ]]
# suppress all but compare
[[ $(./ginkgolinter --suppress-nil-assertion=true --suppress-err-assertion=true --suppress-len-assertion=true --suppress-async-assertion=true --suppress-type-compare-assertion=true a/... 2>&1 | wc -l) == 319 ]]
# suppress all but async
[[ $(./ginkgolinter --suppress-nil-assertion=true --suppress-err-assertion=true --suppress-len-assertion=true --suppress-compare-assertion=true --suppress-type-compare-assertion=true a/... 2>&1 | wc -l) == 183 ]]
# suppress all but focus
[[ $(./ginkgolinter --suppress-nil-assertion=true --suppress-err-assertion=true --suppress-len-assertion=true --suppress-compare-assertion=true --suppress-async-assertion=true --forbid-focus-container=true --suppress-type-compare-assertion=true a/... 2>&1 | wc -l) == 216 ]]
# suppress all but compare different types
[[ $(./ginkgolinter --suppress-nil-assertion=true --suppress-err-assertion=true --suppress-len-assertion=true --suppress-compare-assertion=true --suppress-compare-assertion=true a/... 2>&1 | wc -l) == 267 ]]
# allow HaveLen(0)
[[ $(./ginkgolinter --allow-havelen-0=true a/... 2>&1 | wc -l) == 2578 ]]
# force Expect with To
[[ $(./ginkgolinter --force-expect-to=true a/... 2>&1 | wc -l) == 2773 ]]
# suppress all - should only return the few non-suppressble
[[ $(./ginkgolinter --suppress-nil-assertion=true --suppress-len-assertion=true --suppress-err-assertion=true --suppress-compare-assertion=true --suppress-async-assertion=true --forbid-focus-container=false --suppress-type-compare-assertion=true a/... 2>&1 | wc -l) == 152 ]]
# suppress all, force  Expect with To
[[ $(./ginkgolinter --force-expect-to=true --suppress-nil-assertion=true --suppress-len-assertion=true --suppress-err-assertion=true --suppress-compare-assertion=true --suppress-async-assertion=true --forbid-focus-container=false --suppress-type-compare-assertion=true a/... 2>&1 | wc -l) == 775 ]]
# enable async interval validation
[[ $(./ginkgolinter --validate-async-intervals=true a/... 2>&1 | wc -l) == 2686 ]]
# suppress spec pollution
[[ $(./ginkgolinter --forbid-spec-pollution=true a/... 2>&1 | wc -l) == 2689 ]]
# suppress all but spec pollution
[[ $(./ginkgolinter --suppress-nil-assertion=true --suppress-err-assertion=true --suppress-len-assertion=true --suppress-compare-assertion=true --suppress-async-assertion=true --forbid-spec-pollution=true --suppress-type-compare-assertion=true a/... 2>&1 | wc -l) == 250 ]]
# suppress all but spec pollution && focus containers
[[ $(./ginkgolinter --suppress-nil-assertion=true --suppress-err-assertion=true --suppress-len-assertion=true --suppress-compare-assertion=true --suppress-async-assertion=true --forbid-spec-pollution=true --suppress-type-compare-assertion=true --forbid-focus-container=true a/... 2>&1 | wc -l) == 314 ]]
