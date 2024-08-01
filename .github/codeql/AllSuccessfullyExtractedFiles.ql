/**
 * @id go/diagnostics/all-successfully-extracted-files
 * @name Extracted files
 * @description List all files that were extracted (regardless of relative path).
 * @kind diagnostic
 * @tags successfully-extracted-files
 */

import go

from File f
//where exists(f.getRelativePath())
select f, ""
