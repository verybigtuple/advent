use std::cmp::{max, min};
use std::error;
use std::fmt;
use std::fs::File;
use std::io::{BufRead, BufReader};

#[derive(Default, Clone, Copy)]
struct Dem {
    length: i32,
    width: i32,
    height: i32,
}

#[derive(Debug)]
struct ParseError;

impl fmt::Display for ParseError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "Cannot find the next int")
    }
}

impl error::Error for ParseError {}

fn parse_line(s: &str) -> Result<Dem, Box<dyn error::Error>> {
    let mut split_str = s.split('x');

    let mut d: Dem = Default::default();
    d.length = split_str.next().ok_or(ParseError)?.parse::<i32>()?;
    d.width = split_str.next().ok_or(ParseError)?.parse::<i32>()?;
    d.height = split_str.next().ok_or(ParseError)?.parse::<i32>()?;

    Ok(d)
}

fn calc_wrapper_area(d: Dem) -> i32 {
    let (a, b, c) = (d.length * d.width, d.width * d.height, d.height * d.length);
    let m = min(min(a, b), c);
    2 * a + 2 * b + 2 * c + m
}

fn calc_ribbon(d: Dem) -> i32 {
    let m = max(max(d.height, d.length), d.width);
    let wrap = 2 * d.height + 2 * d.length + 2 * d.width - 2 * m;
    let bow = d.height * d.length * d.width;
    wrap + bow
}

fn main() -> Result<(), Box<dyn error::Error>> {
    let file = File::open("input.txt")?;
    let reader = BufReader::new(file);

    let mut w_area = 0;
    let mut rib = 0;
    for line in reader.lines() {
        let d = parse_line(&line?)?;
        w_area += calc_wrapper_area(d);
        rib += calc_ribbon(d)
    }
    println!("Wrapper area: {}, Ribbon {}", w_area, rib);

    Ok(())
}
