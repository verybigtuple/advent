use std::fs::File;
use std::io::{BufRead, BufReader};

mod bitmap;
mod range;
mod parser;


fn do_instruction(bitmap: &mut bitmap::BitMap, intruct: &str) {
    let mut paser = parser::Parser::new(intruct);
    let command = paser.parse().unwrap();
    let range = range::PointRange::new(command.from, command.to);
    for p in range {
        match command.op {
            parser::Operation::Toggle => bitmap.toggle_bit(p),
            parser::Operation::TurnOff => bitmap.reset_bit(p),
            parser::Operation::TurnOn => bitmap.set_bit(p),
        }
    }
}

fn main() {
    let mut bm = bitmap::BitMap::new(1_000_000);
    let file = File::open("input.txt").unwrap();
    let reader = BufReader::new(file);

    for line in reader.lines() {
        do_instruction(&mut bm, line.unwrap().as_str());
    }
    println!("{}", bm.count())
}


#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test1 () {
        let a = "toggle 0,0 through 999,0";
        let mut bm = bitmap::BitMap::new(1_000_000);
        do_instruction(&mut bm, a);
        assert_eq!(1_000, bm.count());
    }
}