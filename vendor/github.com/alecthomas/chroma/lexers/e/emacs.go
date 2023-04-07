package e

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

var (
	emacsMacros = []string{
		"atomic-change-group", "case", "block", "cl-block", "cl-callf", "cl-callf2",
		"cl-case", "cl-decf", "cl-declaim", "cl-declare",
		"cl-define-compiler-macro", "cl-defmacro", "cl-defstruct",
		"cl-defsubst", "cl-deftype", "cl-defun", "cl-destructuring-bind",
		"cl-do", "cl-do*", "cl-do-all-symbols", "cl-do-symbols", "cl-dolist",
		"cl-dotimes", "cl-ecase", "cl-etypecase", "eval-when", "cl-eval-when", "cl-flet",
		"cl-flet*", "cl-function", "cl-incf", "cl-labels", "cl-letf",
		"cl-letf*", "cl-load-time-value", "cl-locally", "cl-loop",
		"cl-macrolet", "cl-multiple-value-bind", "cl-multiple-value-setq",
		"cl-progv", "cl-psetf", "cl-psetq", "cl-pushnew", "cl-remf",
		"cl-return", "cl-return-from", "cl-rotatef", "cl-shiftf",
		"cl-symbol-macrolet", "cl-tagbody", "cl-the", "cl-typecase",
		"combine-after-change-calls", "condition-case-unless-debug", "decf",
		"declaim", "declare", "declare-function", "def-edebug-spec",
		"defadvice", "defclass", "defcustom", "defface", "defgeneric",
		"defgroup", "define-advice", "define-alternatives",
		"define-compiler-macro", "define-derived-mode", "define-generic-mode",
		"define-global-minor-mode", "define-globalized-minor-mode",
		"define-minor-mode", "define-modify-macro",
		"define-obsolete-face-alias", "define-obsolete-function-alias",
		"define-obsolete-variable-alias", "define-setf-expander",
		"define-skeleton", "defmacro", "defmethod", "defsetf", "defstruct",
		"defsubst", "deftheme", "deftype", "defun", "defvar-local",
		"delay-mode-hooks", "destructuring-bind", "do", "do*",
		"do-all-symbols", "do-symbols", "dolist", "dont-compile", "dotimes",
		"dotimes-with-progress-reporter", "ecase", "ert-deftest", "etypecase",
		"eval-and-compile", "eval-when-compile", "flet", "ignore-errors",
		"incf", "labels", "lambda", "letrec", "lexical-let", "lexical-let*",
		"loop", "multiple-value-bind", "multiple-value-setq", "noreturn",
		"oref", "oref-default", "oset", "oset-default", "pcase",
		"pcase-defmacro", "pcase-dolist", "pcase-exhaustive", "pcase-let",
		"pcase-let*", "pop", "psetf", "psetq", "push", "pushnew", "remf",
		"return", "rotatef", "rx", "save-match-data", "save-selected-window",
		"save-window-excursion", "setf", "setq-local", "shiftf",
		"track-mouse", "typecase", "unless", "use-package", "when",
		"while-no-input", "with-case-table", "with-category-table",
		"with-coding-priority", "with-current-buffer", "with-demoted-errors",
		"with-eval-after-load", "with-file-modes", "with-local-quit",
		"with-output-to-string", "with-output-to-temp-buffer",
		"with-parsed-tramp-file-name", "with-selected-frame",
		"with-selected-window", "with-silent-modifications", "with-slots",
		"with-syntax-table", "with-temp-buffer", "with-temp-file",
		"with-temp-message", "with-timeout", "with-tramp-connection-property",
		"with-tramp-file-property", "with-tramp-progress-reporter",
		"with-wrapper-hook", "load-time-value", "locally", "macrolet", "progv",
		"return-from",
	}

	emacsSpecialForms = []string{
		"and", "catch", "cond", "condition-case", "defconst", "defvar",
		"function", "if", "interactive", "let", "let*", "or", "prog1",
		"prog2", "progn", "quote", "save-current-buffer", "save-excursion",
		"save-restriction", "setq", "setq-default", "subr-arity",
		"unwind-protect", "while",
	}

	emacsBuiltinFunction = []string{
		"%", "*", "+", "-", "/", "/=", "1+", "1-", "<", "<=", "=", ">", ">=",
		"Snarf-documentation", "abort-recursive-edit", "abs",
		"accept-process-output", "access-file", "accessible-keymaps", "acos",
		"active-minibuffer-window", "add-face-text-property",
		"add-name-to-file", "add-text-properties", "all-completions",
		"append", "apply", "apropos-internal", "aref", "arrayp", "aset",
		"ash", "asin", "assoc", "assoc-string", "assq", "atan", "atom",
		"autoload", "autoload-do-load", "backtrace", "backtrace--locals",
		"backtrace-debug", "backtrace-eval", "backtrace-frame",
		"backward-char", "backward-prefix-chars", "barf-if-buffer-read-only",
		"base64-decode-region", "base64-decode-string",
		"base64-encode-region", "base64-encode-string", "beginning-of-line",
		"bidi-find-overridden-directionality", "bidi-resolved-levels",
		"bitmap-spec-p", "bobp", "bolp", "bool-vector",
		"bool-vector-count-consecutive", "bool-vector-count-population",
		"bool-vector-exclusive-or", "bool-vector-intersection",
		"bool-vector-not", "bool-vector-p", "bool-vector-set-difference",
		"bool-vector-subsetp", "bool-vector-union", "boundp",
		"buffer-base-buffer", "buffer-chars-modified-tick",
		"buffer-enable-undo", "buffer-file-name", "buffer-has-markers-at",
		"buffer-list", "buffer-live-p", "buffer-local-value",
		"buffer-local-variables", "buffer-modified-p", "buffer-modified-tick",
		"buffer-name", "buffer-size", "buffer-string", "buffer-substring",
		"buffer-substring-no-properties", "buffer-swap-text", "bufferp",
		"bury-buffer-internal", "byte-code", "byte-code-function-p",
		"byte-to-position", "byte-to-string", "byteorder",
		"call-interactively", "call-last-kbd-macro", "call-process",
		"call-process-region", "cancel-kbd-macro-events", "capitalize",
		"capitalize-region", "capitalize-word", "car", "car-less-than-car",
		"car-safe", "case-table-p", "category-docstring",
		"category-set-mnemonics", "category-table", "category-table-p",
		"ccl-execute", "ccl-execute-on-string", "ccl-program-p", "cdr",
		"cdr-safe", "ceiling", "char-after", "char-before",
		"char-category-set", "char-charset", "char-equal", "char-or-string-p",
		"char-resolve-modifiers", "char-syntax", "char-table-extra-slot",
		"char-table-p", "char-table-parent", "char-table-range",
		"char-table-subtype", "char-to-string", "char-width", "characterp",
		"charset-after", "charset-id-internal", "charset-plist",
		"charset-priority-list", "charsetp", "check-coding-system",
		"check-coding-systems-region", "clear-buffer-auto-save-failure",
		"clear-charset-maps", "clear-face-cache", "clear-font-cache",
		"clear-image-cache", "clear-string", "clear-this-command-keys",
		"close-font", "clrhash", "coding-system-aliases",
		"coding-system-base", "coding-system-eol-type", "coding-system-p",
		"coding-system-plist", "coding-system-priority-list",
		"coding-system-put", "color-distance", "color-gray-p",
		"color-supported-p", "combine-after-change-execute",
		"command-error-default-function", "command-remapping", "commandp",
		"compare-buffer-substrings", "compare-strings",
		"compare-window-configurations", "completing-read",
		"compose-region-internal", "compose-string-internal",
		"composition-get-gstring", "compute-motion", "concat", "cons",
		"consp", "constrain-to-field", "continue-process",
		"controlling-tty-p", "coordinates-in-window-p", "copy-alist",
		"copy-category-table", "copy-file", "copy-hash-table", "copy-keymap",
		"copy-marker", "copy-sequence", "copy-syntax-table", "copysign",
		"cos", "current-active-maps", "current-bidi-paragraph-direction",
		"current-buffer", "current-case-table", "current-column",
		"current-global-map", "current-idle-time", "current-indentation",
		"current-input-mode", "current-local-map", "current-message",
		"current-minor-mode-maps", "current-time", "current-time-string",
		"current-time-zone", "current-window-configuration",
		"cygwin-convert-file-name-from-windows",
		"cygwin-convert-file-name-to-windows", "daemon-initialized",
		"daemonp", "dbus--init-bus", "dbus-get-unique-name",
		"dbus-message-internal", "debug-timer-check", "declare-equiv-charset",
		"decode-big5-char", "decode-char", "decode-coding-region",
		"decode-coding-string", "decode-sjis-char", "decode-time",
		"default-boundp", "default-file-modes", "default-printer-name",
		"default-toplevel-value", "default-value", "define-category",
		"define-charset-alias", "define-charset-internal",
		"define-coding-system-alias", "define-coding-system-internal",
		"define-fringe-bitmap", "define-hash-table-test", "define-key",
		"define-prefix-command", "delete",
		"delete-all-overlays", "delete-and-extract-region", "delete-char",
		"delete-directory-internal", "delete-field", "delete-file",
		"delete-frame", "delete-other-windows-internal", "delete-overlay",
		"delete-process", "delete-region", "delete-terminal",
		"delete-window-internal", "delq", "describe-buffer-bindings",
		"describe-vector", "destroy-fringe-bitmap", "detect-coding-region",
		"detect-coding-string", "ding", "directory-file-name",
		"directory-files", "directory-files-and-attributes", "discard-input",
		"display-supports-face-attributes-p", "do-auto-save", "documentation",
		"documentation-property", "downcase", "downcase-region",
		"downcase-word", "draw-string", "dump-colors", "dump-emacs",
		"dump-face", "dump-frame-glyph-matrix", "dump-glyph-matrix",
		"dump-glyph-row", "dump-redisplay-history", "dump-tool-bar-row",
		"elt", "emacs-pid", "encode-big5-char", "encode-char",
		"encode-coding-region", "encode-coding-string", "encode-sjis-char",
		"encode-time", "end-kbd-macro", "end-of-line", "eobp", "eolp", "eq",
		"eql", "equal", "equal-including-properties", "erase-buffer",
		"error-message-string", "eval", "eval-buffer", "eval-region",
		"event-convert-list", "execute-kbd-macro", "exit-recursive-edit",
		"exp", "expand-file-name", "expt", "external-debugging-output",
		"face-attribute-relative-p", "face-attributes-as-vector", "face-font",
		"fboundp", "fceiling", "fetch-bytecode", "ffloor",
		"field-beginning", "field-end", "field-string",
		"field-string-no-properties", "file-accessible-directory-p",
		"file-acl", "file-attributes", "file-attributes-lessp",
		"file-directory-p", "file-executable-p", "file-exists-p",
		"file-locked-p", "file-modes", "file-name-absolute-p",
		"file-name-all-completions", "file-name-as-directory",
		"file-name-completion", "file-name-directory",
		"file-name-nondirectory", "file-newer-than-file-p", "file-readable-p",
		"file-regular-p", "file-selinux-context", "file-symlink-p",
		"file-system-info", "file-system-info", "file-writable-p",
		"fillarray", "find-charset-region", "find-charset-string",
		"find-coding-systems-region-internal", "find-composition-internal",
		"find-file-name-handler", "find-font", "find-operation-coding-system",
		"float", "float-time", "floatp", "floor", "fmakunbound",
		"following-char", "font-at", "font-drive-otf", "font-face-attributes",
		"font-family-list", "font-get", "font-get-glyphs",
		"font-get-system-font", "font-get-system-normal-font", "font-info",
		"font-match-p", "font-otf-alternates", "font-put",
		"font-shape-gstring", "font-spec", "font-variation-glyphs",
		"font-xlfd-name", "fontp", "fontset-font", "fontset-info",
		"fontset-list", "fontset-list-all", "force-mode-line-update",
		"force-window-update", "format", "format-mode-line",
		"format-network-address", "format-time-string", "forward-char",
		"forward-comment", "forward-line", "forward-word",
		"frame-border-width", "frame-bottom-divider-width",
		"frame-can-run-window-configuration-change-hook", "frame-char-height",
		"frame-char-width", "frame-face-alist", "frame-first-window",
		"frame-focus", "frame-font-cache", "frame-fringe-width", "frame-list",
		"frame-live-p", "frame-or-buffer-changed-p", "frame-parameter",
		"frame-parameters", "frame-pixel-height", "frame-pixel-width",
		"frame-pointer-visible-p", "frame-right-divider-width",
		"frame-root-window", "frame-scroll-bar-height",
		"frame-scroll-bar-width", "frame-selected-window", "frame-terminal",
		"frame-text-cols", "frame-text-height", "frame-text-lines",
		"frame-text-width", "frame-total-cols", "frame-total-lines",
		"frame-visible-p", "framep", "frexp", "fringe-bitmaps-at-pos",
		"fround", "fset", "ftruncate", "funcall", "funcall-interactively",
		"function-equal", "functionp", "gap-position", "gap-size",
		"garbage-collect", "gc-status", "generate-new-buffer-name", "get",
		"get-buffer", "get-buffer-create", "get-buffer-process",
		"get-buffer-window", "get-byte", "get-char-property",
		"get-char-property-and-overlay", "get-file-buffer", "get-file-char",
		"get-internal-run-time", "get-load-suffixes", "get-pos-property",
		"get-process", "get-screen-color", "get-text-property",
		"get-unicode-property-internal", "get-unused-category",
		"get-unused-iso-final-char", "getenv-internal", "gethash",
		"gfile-add-watch", "gfile-rm-watch", "global-key-binding",
		"gnutls-available-p", "gnutls-boot", "gnutls-bye", "gnutls-deinit",
		"gnutls-error-fatalp", "gnutls-error-string", "gnutls-errorp",
		"gnutls-get-initstage", "gnutls-peer-status",
		"gnutls-peer-status-warning-describe", "goto-char", "gpm-mouse-start",
		"gpm-mouse-stop", "group-gid", "group-real-gid",
		"handle-save-session", "handle-switch-frame", "hash-table-count",
		"hash-table-p", "hash-table-rehash-size",
		"hash-table-rehash-threshold", "hash-table-size", "hash-table-test",
		"hash-table-weakness", "iconify-frame", "identity", "image-flush",
		"image-mask-p", "image-metadata", "image-size", "imagemagick-types",
		"imagep", "indent-to", "indirect-function", "indirect-variable",
		"init-image-library", "inotify-add-watch", "inotify-rm-watch",
		"input-pending-p", "insert", "insert-and-inherit",
		"insert-before-markers", "insert-before-markers-and-inherit",
		"insert-buffer-substring", "insert-byte", "insert-char",
		"insert-file-contents", "insert-startup-screen", "int86",
		"integer-or-marker-p", "integerp", "interactive-form", "intern",
		"intern-soft", "internal--track-mouse", "internal-char-font",
		"internal-complete-buffer", "internal-copy-lisp-face",
		"internal-default-process-filter",
		"internal-default-process-sentinel", "internal-describe-syntax-value",
		"internal-event-symbol-parse-modifiers",
		"internal-face-x-get-resource", "internal-get-lisp-face-attribute",
		"internal-lisp-face-attribute-values", "internal-lisp-face-empty-p",
		"internal-lisp-face-equal-p", "internal-lisp-face-p",
		"internal-make-lisp-face", "internal-make-var-non-special",
		"internal-merge-in-global-face",
		"internal-set-alternative-font-family-alist",
		"internal-set-alternative-font-registry-alist",
		"internal-set-font-selection-order",
		"internal-set-lisp-face-attribute",
		"internal-set-lisp-face-attribute-from-resource",
		"internal-show-cursor", "internal-show-cursor-p", "interrupt-process",
		"invisible-p", "invocation-directory", "invocation-name", "isnan",
		"iso-charset", "key-binding", "key-description",
		"keyboard-coding-system", "keymap-parent", "keymap-prompt", "keymapp",
		"keywordp", "kill-all-local-variables", "kill-buffer", "kill-emacs",
		"kill-local-variable", "kill-process", "last-nonminibuffer-frame",
		"lax-plist-get", "lax-plist-put", "ldexp", "length",
		"libxml-parse-html-region", "libxml-parse-xml-region",
		"line-beginning-position", "line-end-position", "line-pixel-height",
		"list", "list-fonts", "list-system-processes", "listp", "load",
		"load-average", "local-key-binding", "local-variable-if-set-p",
		"local-variable-p", "locale-info", "locate-file-internal",
		"lock-buffer", "log", "logand", "logb", "logior", "lognot", "logxor",
		"looking-at", "lookup-image", "lookup-image-map", "lookup-key",
		"lower-frame", "lsh", "macroexpand", "make-bool-vector",
		"make-byte-code", "make-category-set", "make-category-table",
		"make-char", "make-char-table", "make-directory-internal",
		"make-frame-invisible", "make-frame-visible", "make-hash-table",
		"make-indirect-buffer", "make-keymap", "make-list",
		"make-local-variable", "make-marker", "make-network-process",
		"make-overlay", "make-serial-process", "make-sparse-keymap",
		"make-string", "make-symbol", "make-symbolic-link", "make-temp-name",
		"make-terminal-frame", "make-variable-buffer-local",
		"make-variable-frame-local", "make-vector", "makunbound",
		"map-char-table", "map-charset-chars", "map-keymap",
		"map-keymap-internal", "mapatoms", "mapc", "mapcar", "mapconcat",
		"maphash", "mark-marker", "marker-buffer", "marker-insertion-type",
		"marker-position", "markerp", "match-beginning", "match-data",
		"match-end", "matching-paren", "max", "max-char", "md5", "member",
		"memory-info", "memory-limit", "memory-use-counts", "memq", "memql",
		"menu-bar-menu-at-x-y", "menu-or-popup-active-p",
		"menu-or-popup-active-p", "merge-face-attribute", "message",
		"message-box", "message-or-box", "min",
		"minibuffer-completion-contents", "minibuffer-contents",
		"minibuffer-contents-no-properties", "minibuffer-depth",
		"minibuffer-prompt", "minibuffer-prompt-end",
		"minibuffer-selected-window", "minibuffer-window", "minibufferp",
		"minor-mode-key-binding", "mod", "modify-category-entry",
		"modify-frame-parameters", "modify-syntax-entry",
		"mouse-pixel-position", "mouse-position", "move-overlay",
		"move-point-visually", "move-to-column", "move-to-window-line",
		"msdos-downcase-filename", "msdos-long-file-names", "msdos-memget",
		"msdos-memput", "msdos-mouse-disable", "msdos-mouse-enable",
		"msdos-mouse-init", "msdos-mouse-p", "msdos-remember-default-colors",
		"msdos-set-keyboard", "msdos-set-mouse-buttons",
		"multibyte-char-to-unibyte", "multibyte-string-p", "narrow-to-region",
		"natnump", "nconc", "network-interface-info",
		"network-interface-list", "new-fontset", "newline-cache-check",
		"next-char-property-change", "next-frame", "next-overlay-change",
		"next-property-change", "next-read-file-uses-dialog-p",
		"next-single-char-property-change", "next-single-property-change",
		"next-window", "nlistp", "nreverse", "nth", "nthcdr", "null",
		"number-or-marker-p", "number-to-string", "numberp",
		"open-dribble-file", "open-font", "open-termscript",
		"optimize-char-table", "other-buffer", "other-window-for-scrolling",
		"overlay-buffer", "overlay-end", "overlay-get", "overlay-lists",
		"overlay-properties", "overlay-put", "overlay-recenter",
		"overlay-start", "overlayp", "overlays-at", "overlays-in",
		"parse-partial-sexp", "play-sound-internal", "plist-get",
		"plist-member", "plist-put", "point", "point-marker", "point-max",
		"point-max-marker", "point-min", "point-min-marker",
		"pos-visible-in-window-p", "position-bytes", "posix-looking-at",
		"posix-search-backward", "posix-search-forward", "posix-string-match",
		"posn-at-point", "posn-at-x-y", "preceding-char",
		"prefix-numeric-value", "previous-char-property-change",
		"previous-frame", "previous-overlay-change",
		"previous-property-change", "previous-single-char-property-change",
		"previous-single-property-change", "previous-window", "prin1",
		"prin1-to-string", "princ", "print", "process-attributes",
		"process-buffer", "process-coding-system", "process-command",
		"process-connection", "process-contact", "process-datagram-address",
		"process-exit-status", "process-filter", "process-filter-multibyte-p",
		"process-id", "process-inherit-coding-system-flag", "process-list",
		"process-mark", "process-name", "process-plist",
		"process-query-on-exit-flag", "process-running-child-p",
		"process-send-eof", "process-send-region", "process-send-string",
		"process-sentinel", "process-status", "process-tty-name",
		"process-type", "processp", "profiler-cpu-log",
		"profiler-cpu-running-p", "profiler-cpu-start", "profiler-cpu-stop",
		"profiler-memory-log", "profiler-memory-running-p",
		"profiler-memory-start", "profiler-memory-stop", "propertize",
		"purecopy", "put", "put-text-property",
		"put-unicode-property-internal", "puthash", "query-font",
		"query-fontset", "quit-process", "raise-frame", "random", "rassoc",
		"rassq", "re-search-backward", "re-search-forward", "read",
		"read-buffer", "read-char", "read-char-exclusive",
		"read-coding-system", "read-command", "read-event",
		"read-from-minibuffer", "read-from-string", "read-function",
		"read-key-sequence", "read-key-sequence-vector",
		"read-no-blanks-input", "read-non-nil-coding-system", "read-string",
		"read-variable", "recent-auto-save-p", "recent-doskeys",
		"recent-keys", "recenter", "recursion-depth", "recursive-edit",
		"redirect-debugging-output", "redirect-frame-focus", "redisplay",
		"redraw-display", "redraw-frame", "regexp-quote", "region-beginning",
		"region-end", "register-ccl-program", "register-code-conversion-map",
		"remhash", "remove-list-of-text-properties", "remove-text-properties",
		"rename-buffer", "rename-file", "replace-match",
		"reset-this-command-lengths", "resize-mini-window-internal",
		"restore-buffer-modified-p", "resume-tty", "reverse", "round",
		"run-hook-with-args", "run-hook-with-args-until-failure",
		"run-hook-with-args-until-success", "run-hook-wrapped", "run-hooks",
		"run-window-configuration-change-hook", "run-window-scroll-functions",
		"safe-length", "scan-lists", "scan-sexps", "scroll-down",
		"scroll-left", "scroll-other-window", "scroll-right", "scroll-up",
		"search-backward", "search-forward", "secure-hash", "select-frame",
		"select-window", "selected-frame", "selected-window",
		"self-insert-command", "send-string-to-terminal", "sequencep",
		"serial-process-configure", "set", "set-buffer",
		"set-buffer-auto-saved", "set-buffer-major-mode",
		"set-buffer-modified-p", "set-buffer-multibyte", "set-case-table",
		"set-category-table", "set-char-table-extra-slot",
		"set-char-table-parent", "set-char-table-range", "set-charset-plist",
		"set-charset-priority", "set-coding-system-priority",
		"set-cursor-size", "set-default", "set-default-file-modes",
		"set-default-toplevel-value", "set-file-acl", "set-file-modes",
		"set-file-selinux-context", "set-file-times", "set-fontset-font",
		"set-frame-height", "set-frame-position", "set-frame-selected-window",
		"set-frame-size", "set-frame-width", "set-fringe-bitmap-face",
		"set-input-interrupt-mode", "set-input-meta-mode", "set-input-mode",
		"set-keyboard-coding-system-internal", "set-keymap-parent",
		"set-marker", "set-marker-insertion-type", "set-match-data",
		"set-message-beep", "set-minibuffer-window",
		"set-mouse-pixel-position", "set-mouse-position",
		"set-network-process-option", "set-output-flow-control",
		"set-process-buffer", "set-process-coding-system",
		"set-process-datagram-address", "set-process-filter",
		"set-process-filter-multibyte",
		"set-process-inherit-coding-system-flag", "set-process-plist",
		"set-process-query-on-exit-flag", "set-process-sentinel",
		"set-process-window-size", "set-quit-char",
		"set-safe-terminal-coding-system-internal", "set-screen-color",
		"set-standard-case-table", "set-syntax-table",
		"set-terminal-coding-system-internal", "set-terminal-local-value",
		"set-terminal-parameter", "set-text-properties", "set-time-zone-rule",
		"set-visited-file-modtime", "set-window-buffer",
		"set-window-combination-limit", "set-window-configuration",
		"set-window-dedicated-p", "set-window-display-table",
		"set-window-fringes", "set-window-hscroll", "set-window-margins",
		"set-window-new-normal", "set-window-new-pixel",
		"set-window-new-total", "set-window-next-buffers",
		"set-window-parameter", "set-window-point", "set-window-prev-buffers",
		"set-window-redisplay-end-trigger", "set-window-scroll-bars",
		"set-window-start", "set-window-vscroll", "setcar", "setcdr",
		"setplist", "show-face-resources", "signal", "signal-process", "sin",
		"single-key-description", "skip-chars-backward", "skip-chars-forward",
		"skip-syntax-backward", "skip-syntax-forward", "sleep-for", "sort",
		"sort-charsets", "special-variable-p", "split-char",
		"split-window-internal", "sqrt", "standard-case-table",
		"standard-category-table", "standard-syntax-table", "start-kbd-macro",
		"start-process", "stop-process", "store-kbd-macro-event", "string",
		"string-as-multibyte", "string-as-unibyte", "string-bytes",
		"string-collate-equalp", "string-collate-lessp", "string-equal",
		"string-lessp", "string-make-multibyte", "string-make-unibyte",
		"string-match", "string-to-char", "string-to-multibyte",
		"string-to-number", "string-to-syntax", "string-to-unibyte",
		"string-width", "stringp", "subr-name", "subrp",
		"subst-char-in-region", "substitute-command-keys",
		"substitute-in-file-name", "substring", "substring-no-properties",
		"suspend-emacs", "suspend-tty", "suspicious-object", "sxhash",
		"symbol-function", "symbol-name", "symbol-plist", "symbol-value",
		"symbolp", "syntax-table", "syntax-table-p", "system-groups",
		"system-move-file-to-trash", "system-name", "system-users", "tan",
		"terminal-coding-system", "terminal-list", "terminal-live-p",
		"terminal-local-value", "terminal-name", "terminal-parameter",
		"terminal-parameters", "terpri", "test-completion",
		"text-char-description", "text-properties-at", "text-property-any",
		"text-property-not-all", "this-command-keys",
		"this-command-keys-vector", "this-single-command-keys",
		"this-single-command-raw-keys", "time-add", "time-less-p",
		"time-subtract", "tool-bar-get-system-style", "tool-bar-height",
		"tool-bar-pixel-width", "top-level", "trace-redisplay",
		"trace-to-stderr", "translate-region-internal", "transpose-regions",
		"truncate", "try-completion", "tty-display-color-cells",
		"tty-display-color-p", "tty-no-underline",
		"tty-suppress-bold-inverse-default-colors", "tty-top-frame",
		"tty-type", "type-of", "undo-boundary", "unencodable-char-position",
		"unhandled-file-name-directory", "unibyte-char-to-multibyte",
		"unibyte-string", "unicode-property-table-internal", "unify-charset",
		"unintern", "unix-sync", "unlock-buffer", "upcase", "upcase-initials",
		"upcase-initials-region", "upcase-region", "upcase-word",
		"use-global-map", "use-local-map", "user-full-name",
		"user-login-name", "user-real-login-name", "user-real-uid",
		"user-uid", "variable-binding-locus", "vconcat", "vector",
		"vector-or-char-table-p", "vectorp", "verify-visited-file-modtime",
		"vertical-motion", "visible-frame-list", "visited-file-modtime",
		"w16-get-clipboard-data", "w16-selection-exists-p",
		"w16-set-clipboard-data", "w32-battery-status",
		"w32-default-color-map", "w32-define-rgb-color",
		"w32-display-monitor-attributes-list", "w32-frame-menu-bar-size",
		"w32-frame-rect", "w32-get-clipboard-data",
		"w32-get-codepage-charset", "w32-get-console-codepage",
		"w32-get-console-output-codepage", "w32-get-current-locale-id",
		"w32-get-default-locale-id", "w32-get-keyboard-layout",
		"w32-get-locale-info", "w32-get-valid-codepages",
		"w32-get-valid-keyboard-layouts", "w32-get-valid-locale-ids",
		"w32-has-winsock", "w32-long-file-name", "w32-reconstruct-hot-key",
		"w32-register-hot-key", "w32-registered-hot-keys",
		"w32-selection-exists-p", "w32-send-sys-command",
		"w32-set-clipboard-data", "w32-set-console-codepage",
		"w32-set-console-output-codepage", "w32-set-current-locale",
		"w32-set-keyboard-layout", "w32-set-process-priority",
		"w32-shell-execute", "w32-short-file-name", "w32-toggle-lock-key",
		"w32-unload-winsock", "w32-unregister-hot-key", "w32-window-exists-p",
		"w32notify-add-watch", "w32notify-rm-watch",
		"waiting-for-user-input-p", "where-is-internal", "widen",
		"widget-apply", "widget-get", "widget-put",
		"window-absolute-pixel-edges", "window-at", "window-body-height",
		"window-body-width", "window-bottom-divider-width", "window-buffer",
		"window-combination-limit", "window-configuration-frame",
		"window-configuration-p", "window-dedicated-p",
		"window-display-table", "window-edges", "window-end", "window-frame",
		"window-fringes", "window-header-line-height", "window-hscroll",
		"window-inside-absolute-pixel-edges", "window-inside-edges",
		"window-inside-pixel-edges", "window-left-child",
		"window-left-column", "window-line-height", "window-list",
		"window-list-1", "window-live-p", "window-margins",
		"window-minibuffer-p", "window-mode-line-height", "window-new-normal",
		"window-new-pixel", "window-new-total", "window-next-buffers",
		"window-next-sibling", "window-normal-size", "window-old-point",
		"window-parameter", "window-parameters", "window-parent",
		"window-pixel-edges", "window-pixel-height", "window-pixel-left",
		"window-pixel-top", "window-pixel-width", "window-point",
		"window-prev-buffers", "window-prev-sibling",
		"window-redisplay-end-trigger", "window-resize-apply",
		"window-resize-apply-total", "window-right-divider-width",
		"window-scroll-bar-height", "window-scroll-bar-width",
		"window-scroll-bars", "window-start", "window-system",
		"window-text-height", "window-text-pixel-size", "window-text-width",
		"window-top-child", "window-top-line", "window-total-height",
		"window-total-width", "window-use-time", "window-valid-p",
		"window-vscroll", "windowp", "write-char", "write-region",
		"x-backspace-delete-keys-p", "x-change-window-property",
		"x-change-window-property", "x-close-connection",
		"x-close-connection", "x-create-frame", "x-create-frame",
		"x-delete-window-property", "x-delete-window-property",
		"x-disown-selection-internal", "x-display-backing-store",
		"x-display-backing-store", "x-display-color-cells",
		"x-display-color-cells", "x-display-grayscale-p",
		"x-display-grayscale-p", "x-display-list", "x-display-list",
		"x-display-mm-height", "x-display-mm-height", "x-display-mm-width",
		"x-display-mm-width", "x-display-monitor-attributes-list",
		"x-display-pixel-height", "x-display-pixel-height",
		"x-display-pixel-width", "x-display-pixel-width", "x-display-planes",
		"x-display-planes", "x-display-save-under", "x-display-save-under",
		"x-display-screens", "x-display-screens", "x-display-visual-class",
		"x-display-visual-class", "x-family-fonts", "x-file-dialog",
		"x-file-dialog", "x-file-dialog", "x-focus-frame", "x-frame-geometry",
		"x-frame-geometry", "x-get-atom-name", "x-get-resource",
		"x-get-selection-internal", "x-hide-tip", "x-hide-tip",
		"x-list-fonts", "x-load-color-file", "x-menu-bar-open-internal",
		"x-menu-bar-open-internal", "x-open-connection", "x-open-connection",
		"x-own-selection-internal", "x-parse-geometry", "x-popup-dialog",
		"x-popup-menu", "x-register-dnd-atom", "x-select-font",
		"x-select-font", "x-selection-exists-p", "x-selection-owner-p",
		"x-send-client-message", "x-server-max-request-size",
		"x-server-max-request-size", "x-server-vendor", "x-server-vendor",
		"x-server-version", "x-server-version", "x-show-tip", "x-show-tip",
		"x-synchronize", "x-synchronize", "x-uses-old-gtk-dialog",
		"x-window-property", "x-window-property", "x-wm-set-size-hint",
		"xw-color-defined-p", "xw-color-defined-p", "xw-color-values",
		"xw-color-values", "xw-display-color-p", "xw-display-color-p",
		"yes-or-no-p", "zlib-available-p", "zlib-decompress-region",
		"forward-point",
	}

	emacsBuiltinFunctionHighlighted = []string{
		"defvaralias", "provide", "require",
		"with-no-warnings", "define-widget", "with-electric-help",
		"throw", "defalias", "featurep",
	}

	emacsLambdaListKeywords = []string{
		"&allow-other-keys", "&aux", "&body", "&environment", "&key", "&optional",
		"&rest", "&whole",
	}

	emacsErrorKeywords = []string{
		"cl-assert", "cl-check-type", "error", "signal",
		"user-error", "warn",
	}
)

// EmacsLisp lexer.
var EmacsLisp = internal.Register(TypeRemappingLexer(MustNewLazyLexer(
	&Config{
		Name:      "EmacsLisp",
		Aliases:   []string{"emacs", "elisp", "emacs-lisp"},
		Filenames: []string{"*.el"},
		MimeTypes: []string{"text/x-elisp", "application/x-elisp"},
	},
	emacsLispRules,
), TypeMapping{
	{NameVariable, NameFunction, emacsBuiltinFunction},
	{NameVariable, NameBuiltin, emacsSpecialForms},
	{NameVariable, NameException, emacsErrorKeywords},
	{NameVariable, NameBuiltin, append(emacsBuiltinFunctionHighlighted, emacsMacros...)},
	{NameVariable, KeywordPseudo, emacsLambdaListKeywords},
}))

func emacsLispRules() Rules {
	return Rules{
		"root": {
			Default(Push("body")),
		},
		"body": {
			{`\s+`, Text, nil},
			{`;.*$`, CommentSingle, nil},
			{`"`, LiteralString, Push("string")},
			{`\?([^\\]|\\.)`, LiteralStringChar, nil},
			{`:((?:\\.|[\w!$%&*+-/<=>?@^{}~|])(?:\\.|[\w!$%&*+-/<=>?@^{}~|]|[#.:])*)`, NameBuiltin, nil},
			{`::((?:\\.|[\w!$%&*+-/<=>?@^{}~|])(?:\\.|[\w!$%&*+-/<=>?@^{}~|]|[#.:])*)`, LiteralStringSymbol, nil},
			{`'((?:\\.|[\w!$%&*+-/<=>?@^{}~|])(?:\\.|[\w!$%&*+-/<=>?@^{}~|]|[#.:])*)`, LiteralStringSymbol, nil},
			{`'`, Operator, nil},
			{"`", Operator, nil},
			{"[-+]?\\d+\\.?(?=[ \"()\\]\\'\\n,;`])", LiteralNumberInteger, nil},
			{"[-+]?\\d+/\\d+(?=[ \"()\\]\\'\\n,;`])", LiteralNumber, nil},
			{"[-+]?(\\d*\\.\\d+([defls][-+]?\\d+)?|\\d+(\\.\\d*)?[defls][-+]?\\d+)(?=[ \"()\\]\\'\\n,;`])", LiteralNumberFloat, nil},
			{`\[|\]`, Punctuation, nil},
			{`#:((?:\\.|[\w!$%&*+-/<=>?@^{}~|])(?:\\.|[\w!$%&*+-/<=>?@^{}~|]|[#.:])*)`, LiteralStringSymbol, nil},
			{`#\^\^?`, Operator, nil},
			{`#\'`, NameFunction, nil},
			{`#[bB][+-]?[01]+(/[01]+)?`, LiteralNumberBin, nil},
			{`#[oO][+-]?[0-7]+(/[0-7]+)?`, LiteralNumberOct, nil},
			{`#[xX][+-]?[0-9a-fA-F]+(/[0-9a-fA-F]+)?`, LiteralNumberHex, nil},
			{`#\d+r[+-]?[0-9a-zA-Z]+(/[0-9a-zA-Z]+)?`, LiteralNumber, nil},
			{`#\d+=`, Operator, nil},
			{`#\d+#`, Operator, nil},
			{`(,@|,|\.|:)`, Operator, nil},
			{"(t|nil)(?=[ \"()\\]\\'\\n,;`])", NameConstant, nil},
			{`\*((?:\\.|[\w!$%&*+-/<=>?@^{}~|])(?:\\.|[\w!$%&*+-/<=>?@^{}~|]|[#.:])*)\*`, NameVariableGlobal, nil},
			{`((?:\\.|[\w!$%&*+-/<=>?@^{}~|])(?:\\.|[\w!$%&*+-/<=>?@^{}~|]|[#.:])*)`, NameVariable, nil},
			{`#\(`, Operator, Push("body")},
			{`\(`, Punctuation, Push("body")},
			{`\)`, Punctuation, Pop(1)},
		},
		"string": {
			{"[^\"\\\\`]+", LiteralString, nil},
			{"`((?:\\\\.|[\\w!$%&*+-/<=>?@^{}~|])(?:\\\\.|[\\w!$%&*+-/<=>?@^{}~|]|[#.:])*)\\'", LiteralStringSymbol, nil},
			{"`", LiteralString, nil},
			{`\\.`, LiteralString, nil},
			{`\\\n`, LiteralString, nil},
			{`"`, LiteralString, Pop(1)},
		},
	}
}
