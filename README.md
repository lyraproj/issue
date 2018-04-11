# Formalized Warning and Error reporting

This package contains the functionality needed to formalize how warnings
and errors that occur at runtime are reported.

A reported error (represented by the interface issue.Reported) contains
a Code that represents an Issue. An issue has a symbolic name and
a format string and the Reported contains the arguments needed to
produce the complete string, including the location where the issue
was reported.
