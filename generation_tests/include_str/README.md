# include_str generation test

Tests that `include_str!()` and `include_bytes!()` macro invocations are detected
and the referenced files are added to `compile_data` in the generated BUILD file.
