const DATA: &str = include_str!("example.json");

const BINARY: &[u8] = include_bytes!("data/binary.bin");

fn get_data() -> &'static str {
    DATA
}

fn get_binary() -> &'static [u8] {
    BINARY
}
