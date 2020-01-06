# ------------------------------------------------------------
# Copyright (c) Microsoft Corporation.
# Licensed under the MIT License.
# ------------------------------------------------------------

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
CODEGEN_PKG=${CODEGEN_PKG:-$(cd "${SCRIPT_ROOT}"; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo ../code-generator)}

bash ".${CODEGEN_PKG}"/generate-groups.sh "deepcopy,client,informer,lister" \
  "dapr/pkg/client" "dapr/pkg/apis" \
  "components:v1alpha1 configuration:v1alpha1"