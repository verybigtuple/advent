use core::fmt;
use std::collections::HashSet;
use std::error;
use std::fs;
use std::io;

/*
This may seem too much for the problem solution. But I wanted to try:
1) Common error type w/o Boxes
2) Structs and traits
 */

#[derive(Debug)]
enum AppError {
    Parse(ShiftError),
    Io(io::Error),
}

impl error::Error for AppError {
    fn cause(&self) -> Option<&dyn error::Error> {
        match *self {
            AppError::Io(ref err) => Some(err),
            AppError::Parse(ref err) => Some(err),
        }
    }
}

impl fmt::Display for AppError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match *self {
            AppError::Io(ref err) => write!(f, "File error {}", err),
            AppError::Parse(ref err) => write!(f, "Parsing error {}", err),
        }
    }
}

impl From<ShiftError> for AppError {
    fn from(value: ShiftError) -> Self {
        AppError::Parse(value)
    }
}

impl From<io::Error> for AppError {
    fn from(value: io::Error) -> Self {
        AppError::Io(value)
    }
}

#[derive(Debug, Clone, Copy)]
struct ShiftError {
    wrong_char: u8,
}

impl error::Error for ShiftError {}

impl fmt::Display for ShiftError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "Unknown direction: {}", self.wrong_char)
    }
}

#[derive(Clone, Copy, PartialEq, Eq, Hash)]
struct Point(i32, i32);

fn shift_point(p: Point, dir: u8) -> Result<Point, ShiftError> {
    match dir {
        b'>' => Ok(Point(p.0 + 1, p.1)),
        b'<' => Ok(Point(p.0 - 1, p.1)),
        b'^' => Ok(Point(p.0, p.1 + 1)),
        b'v' => Ok(Point(p.0, p.1 - 1)),
        _ => Err(ShiftError { wrong_char: dir }),
    }
}

trait Path {
    fn get_visited(&self) -> &InnerSantaVisited;
    fn move_forward(&mut self, dir: u8) -> Result<(), ShiftError>;
    fn len(&self) -> usize {
        self.get_visited().visited.len()
    }
}

struct InnerSantaVisited {
    visited: HashSet<Point>,
}

struct SantaPath {
    current: Point,
    inner: InnerSantaVisited,
}

impl SantaPath {
    fn new() -> Self {
        let mut s = SantaPath {
            current: Point(0, 0),
            inner: InnerSantaVisited {
                visited: HashSet::new(),
            },
        };
        s.inner.visited.insert(Point(0, 0));
        s
    }
}

impl Path for SantaPath {
    fn get_visited(&self) -> &InnerSantaVisited {
        &self.inner
    }

    fn move_forward(&mut self, dir: u8) -> Result<(), ShiftError> {
        self.current = shift_point(self.current, dir)?;
        self.inner.visited.insert(self.current);
        Ok(())
    }
}

struct SantaRoboPath {
    current_santa: Point,
    current_robo: Point,
    inner: InnerSantaVisited,
    santa_move: bool,
}

impl SantaRoboPath {
    fn new() -> Self {
        let mut s = SantaRoboPath {
            current_robo: Point(0, 0),
            current_santa: Point(0, 0),
            santa_move: true,
            inner: InnerSantaVisited {
                visited: HashSet::new(),
            },
        };
        s.inner.visited.insert(Point(0, 0));
        s
    }
}

impl Path for SantaRoboPath {
    fn get_visited(&self) -> &InnerSantaVisited {
        &self.inner
    }

    fn move_forward(&mut self, dir: u8) -> Result<(), ShiftError> {
        if self.santa_move {
            self.current_santa = shift_point(self.current_santa, dir)?;
            self.inner.visited.insert(self.current_santa);
        } else {
            self.current_robo = shift_point(self.current_robo, dir)?;
            self.inner.visited.insert(self.current_robo);
        }
        self.santa_move = !self.santa_move;

        Ok(())
    }
}

fn process(reader: impl io::Read) -> Result<(), AppError> {
    let mut sp = SantaPath::new();
    let mut srp = SantaRoboPath::new();
    for b in reader.bytes() {
        let byte = b?;
        sp.move_forward(byte)?;
        srp.move_forward(byte)?;
    }
    println!("Santa moved: {}", sp.len());
    println!("Santa+Robo moved: {}", srp.len());
    Ok(())
}

fn main() -> Result<(), AppError> {
    let file = fs::File::open("input.txt")?;
    process(file)?;

    Ok(())
}
