#!/usr/bin/env bash

set -ex

cp ginkgolinter testdata/src/a
cd testdata/src/a

# no suppress
[[ $(./ginkgolinter a/... 2>&1 | wc -l) == 2622 ]]
# suppress all but nil
[[ $(./ginkgolinter --suppress-len-assertion=true --suppress-err-assertion=true --suppress-compare-assertion=true --suppress-async-assertion=true --suppress-type-compare-assertion=true a/... 2>&1 | wc -l) == 1573 ]]
# suppress all but len
[[ $(./ginkgolinter --suppress-nil-assertion=true --suppress-err-assertion=true --suppress-compare-assertion=true --suppress-async-assertion=true --suppress-type-compare-assertion=true a/... 2>&1 | wc -l) == 933 ]]
# suppress all but err
[[ $(./ginkgolinter --suppress-nil-assertion=true --suppress-len-assertion=true --suppress-compare-assertion=true --suppress-async-assertion=true --suppress-type-compare-assertion=true a/... 2>&1 | wc -l) == 303 ]]
# suppress all but compare
[[ $(./ginkgolinter --suppress-nil-assertion=true --suppress-err-assertion=true --suppress-len-assertion=true --suppress-async-assertion=true --suppress-type-compare-assertion=true a/... 2>&1 | wc -l) == 376 ]]
# suppress all but async
[[ $(./ginkgolinter --suppress-nil-assertion=true --suppress-err-assertion=true --suppress-len-assertion=true --suppress-compare-assertion=true --suppress-type-compare-assertion=true a/... 2>&1 | wc -l) == 237 ]]
# suppress all but focus
[[ $(./ginkgolinter --suppress-nil-assertion=true --suppress-err-assertion=true --suppress-len-assertion=true --suppress-compare-assertion=true --suppress-async-assertion=true --forbid-focus-container=true --suppress-type-compare-assertion=true a/... 2>&1 | wc -l) == 270 ]]
# suppress all but compare different types
[[ $(./ginkgolinter --suppress-nil-assertion=true --suppress-err-assertion=true --suppress-len-assertion=true --suppress-compare-assertion=true --suppress-compare-assertion=true a/... 2>&1 | wc -l) == 321 ]]
# allow HaveLen(0)
[[ $(./ginkgolinter --allow-havelen-0=true a/... 2>&1 | wc -l) == 2609 ]]
# force Expect with To
[[ $(./ginkgolinter --force-expect-to=true a/... 2>&1 | wc -l) == 2815 ]]
# suppress all - should only return the few non-suppressble
[[ $(./ginkgolinter --suppress-nil-assertion=true --suppress-len-assertion=true --suppress-err-assertion=true --suppress-compare-assertion=true --suppress-async-assertion=true --forbid-focus-container=false --suppress-type-compare-assertion=true a/... 2>&1 | wc -l) == 206 ]]
# suppress all, force  Expect with To
[[ $(./ginkgolinter --force-expect-to=true --suppress-nil-assertion=true --suppress-len-assertion=true --suppress-err-assertion=true --suppress-compare-assertion=true --suppress-async-assertion=true --forbid-focus-container=false --suppress-type-compare-assertion=true a/... 2>&1 | wc -l) == 1113 ]]
# enable async interval validation
[[ $(./ginkgolinter --validate-async-intervals=true a/... 2>&1 | wc -l) == 2717 ]]
# suppress spec pollution
[[ $(./ginkgolinter --forbid-spec-pollution=true a/... 2>&1 | wc -l) == 2720 ]]
# suppress all but spec pollution
[[ $(./ginkgolinter --suppress-nil-assertion=true --suppress-err-assertion=true --suppress-len-assertion=true --suppress-compare-assertion=true --suppress-async-assertion=true --forbid-spec-pollution=true --suppress-type-compare-assertion=true a/... 2>&1 | wc -l) == 304 ]]
# suppress all but spec pollution && focus containers
[[ $(./ginkgolinter --suppress-nil-assertion=true --suppress-err-assertion=true --suppress-len-assertion=true --suppress-compare-assertion=true --suppress-async-assertion=true --forbid-spec-pollution=true --suppress-type-compare-assertion=true --forbid-focus-container=true a/... 2>&1 | wc -l) == 368 ]]
