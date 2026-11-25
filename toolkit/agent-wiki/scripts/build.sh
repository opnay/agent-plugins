#!/bin/sh
set -eu

script_dir=$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)
cargo build --locked --release --manifest-path "$script_dir/Cargo.toml" --target-dir "$script_dir/target"
install -m 755 "$script_dir/target/release/wdown" "$script_dir/wdown"
printf '%s\n' "$script_dir/wdown"
