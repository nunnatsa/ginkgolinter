#!/usr/bin/env bash

set -ex

cp ginkgolinter testdata/src/a
cd testdata/src/a

# no suppress
[[ $(./ginkgolinter a/... 2>&1 | wc -l) == 2665 ]]
# suppress all but nil
[[ $(./ginkgolinter --suppress-len-assertion=true --suppress-err-assertion=true --suppress-compare-assertion=true --suppress-async-assertion=true --suppress-type-compare-assertion=true --suppress-succeed-assertion=true a/... 2>&1 | wc -l) == 1567 ]]
# suppress all but len
[[ $(./ginkgolinter --suppress-nil-assertion=true --suppress-err-assertion=true --suppress-compare-assertion=true --suppress-async-assertion=true --suppress-type-compare-assertion=true --suppress-succeed-assertion=true a/... 2>&1 | wc -l) == 924 ]]
# suppress all but err
[[ $(./ginkgolinter --suppress-nil-assertion=true --suppress-len-assertion=true --suppress-compare-assertion=true --suppress-async-assertion=true --suppress-type-compare-assertion=true --suppress-succeed-assertion=true a/... 2>&1 | wc -l) == 332 ]]
# suppress all but compare
[[ $(./ginkgolinter --suppress-nil-assertion=true --suppress-err-assertion=true --suppress-len-assertion=true --suppress-async-assertion=true --suppress-type-compare-assertion=true --suppress-succeed-assertion=true a/... 2>&1 | wc -l) == 367 ]]
# suppress all but async
[[ $(./ginkgolinter --suppress-nil-assertion=true --suppress-err-assertion=true --suppress-len-assertion=true --suppress-compare-assertion=true --suppress-type-compare-assertion=true --suppress-succeed-assertion=true a/... 2>&1 | wc -l) == 230 ]]
# suppress all but focus
[[ $(./ginkgolinter --suppress-nil-assertion=true --suppress-err-assertion=true --suppress-len-assertion=true --suppress-compare-assertion=true --suppress-async-assertion=true --forbid-focus-container=true --suppress-type-compare-assertion=true --suppress-succeed-assertion=true a/... 2>&1 | wc -l) == 264 ]]
# suppress all but compare different types
[[ $(./ginkgolinter --suppress-nil-assertion=true --suppress-err-assertion=true --suppress-len-assertion=true --suppress-compare-assertion=true --suppress-compare-assertion=true --suppress-succeed-assertion=true a/... 2>&1 | wc -l) == 314 ]]
# allow HaveLen(0
[[ $(./ginkgolinter --allow-havelen-0=true a/... 2>&1 | wc -l) == 2652 ]]
# force Expect with To
[[ $(./ginkgolinter --force-expect-to=true a/... 2>&1 | wc -l) == 2847 ]]
# suppress all - should only return the few non-suppressble
[[ $(./ginkgolinter --suppress-nil-assertion=true --suppress-len-assertion=true --suppress-err-assertion=true --suppress-compare-assertion=true --suppress-async-assertion=true --forbid-focus-container=false --suppress-type-compare-assertion=true --suppress-succeed-assertion=true a/... 2>&1 | wc -l) == 200 ]]
# suppress all, force  Expect with To
[[ $(./ginkgolinter --force-expect-to=true --suppress-nil-assertion=true --suppress-len-assertion=true --suppress-err-assertion=true --suppress-compare-assertion=true --suppress-async-assertion=true --forbid-focus-container=false --suppress-type-compare-assertion=true --suppress-succeed-assertion=true a/... 2>&1 | wc -l) == 823 ]]
# enable async interval validation
[[ $(./ginkgolinter --validate-async-intervals=true a/... 2>&1 | wc -l) == 2760 ]]
# suppress spec pollution
[[ $(./ginkgolinter --forbid-spec-pollution=true a/... 2>&1 | wc -l) == 2775 ]]
# suppress all but spec pollution
[[ $(./ginkgolinter --suppress-nil-assertion=true --suppress-err-assertion=true --suppress-len-assertion=true --suppress-compare-assertion=true --suppress-async-assertion=true --forbid-spec-pollution=true --suppress-type-compare-assertion=true --suppress-succeed-assertion=true a/... 2>&1 | wc -l) == 310 ]]
# suppress all but spec pollution && focus containers
[[ $(./ginkgolinter --suppress-nil-assertion=true --suppress-err-assertion=true --suppress-len-assertion=true --suppress-compare-assertion=true --suppress-async-assertion=true --forbid-spec-pollution=true --suppress-type-compare-assertion=true --forbid-focus-container=true --suppress-succeed-assertion=true a/... 2>&1 | wc -l) == 374 ]]
# suppress all but succeed 
[[ $(./ginkgolinter --suppress-nil-assertion=true --suppress-len-assertion=true --suppress-err-assertion=true --suppress-compare-assertion=true --suppress-async-assertion=true --forbid-focus-container=false --suppress-type-compare-assertion=true a/... 2>&1 | wc -l) == 216 ]]