use std::fs::File;
use std::io::{BufRead, BufReader};

mod bitmap;
mod bytemap;
mod parser;
mod range;

fn do_onoff(bitmap: &mut bitmap::BitMap, range: range::PointRange, operation: parser::Operation) {
    for p in range {
        match operation {
            parser::Operation::Toggle => bitmap.toggle_bit(p),
            parser::Operation::TurnOff => bitmap.reset_bit(p),
            parser::Operation::TurnOn => bitmap.set_bit(p),
        }
    }
}

fn do_brightness(
    brightness: &mut bytemap::ByteMap,
    range: range::PointRange,
    operation: parser::Operation,
) {
    for p in range {
        match operation {
            parser::Operation::Toggle => brightness.inc_byte_by(p, 2),
            parser::Operation::TurnOff => brightness.dec_byte(p),
            parser::Operation::TurnOn => brightness.inc_byte(p),
        }
    }
}

fn main() {
    let mut bm = bitmap::BitMap::new(1_000_000);
    let mut brightness = bytemap::ByteMap::new();
    let file = File::open("input.txt").unwrap();
    let reader = BufReader::new(file);
    for line in reader.lines() {
        let line = line.unwrap();
        let mut parser = parser::Parser::new(&line);
        let command = parser.parse().unwrap();
        let range = range::PointRange::new(command.from, command.to);
        do_onoff(&mut bm, range, command.op);
        do_brightness(&mut brightness, range, command.op);
    }
    println!("first part: {}", bm.count());
    println!("second part: {}", brightness.count());
}
