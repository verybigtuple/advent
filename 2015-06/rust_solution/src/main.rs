use core::fmt;
use std::error;
use std::fs::File;
use std::io;
use std::io::{BufRead, BufReader};

mod bitmap;
mod bytemap;
mod parser;
mod range;

const LAMP_COUNT: usize = 1000_1000;
const MAX_POINT: range::Point = range::Point(999, 999);

#[derive(Debug)]
enum AppError {
    ParseError(parser::ParseError),
    FileError(io::Error),
}

impl fmt::Display for AppError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match *self {
            AppError::ParseError(ref err) => write!(f, "Parsing line error: {}", err),
            AppError::FileError(ref err) => write!(f, "File error: {}", err),
        }
    }
}

impl error::Error for AppError {
    fn cause(&self) -> Option<&dyn error::Error> {
        match *self {
            AppError::ParseError(ref err) => Some(err),
            AppError::FileError(ref err) => Some(err),
        }
    }
}

impl From<parser::ParseError> for AppError {
    fn from(value: parser::ParseError) -> Self {
        AppError::ParseError(value)
    }
}

impl From<io::Error> for AppError {
    fn from(value: io::Error) -> Self {
        AppError::FileError(value)
    }
}

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

fn main() -> Result<(), AppError> {
    let file = File::open("input.txt")?;

    let mut toggler = bitmap::BitMap::new(LAMP_COUNT);
    let mut brightness = bytemap::ByteMap::new(LAMP_COUNT);

    let reader = BufReader::new(file);
    for line in reader.lines() {
        let line = line?;
        let mut parser = parser::Parser::new(&line);
        let command = parser.parse()?;
        let range = range::PointRange::new(command.from, command.to, MAX_POINT);
        do_onoff(&mut toggler, range, command.op);
        do_brightness(&mut brightness, range, command.op);
    }
    println!("first part: {}", toggler.count());
    println!("second part: {}", brightness.count());
    Ok(())
}
