use std::{
    fs::File,
    io::{Error, Read, Seek},
};

const LEFT_BRACE: u8 = 0x28;

fn solve_floor(reader: &mut impl Read) -> Result<i64, Error> {
    let mut i: i64 = 0;
    for byte in reader.bytes() {
        match byte {
            Ok(b) => i += if b == LEFT_BRACE { 1 } else { -1 },
            Err(e) => return Err(e),
        }
    }
    Ok(i)
}

fn solve_basement(reader: &mut impl Read) -> Result<usize, Error> {
    let mut i: isize = 0;

    for (c, byte) in reader.bytes().enumerate() {
        match byte {
            Ok(b) => {
                i += if b == LEFT_BRACE { 1 } else { -1 };
                if i == -1 {
                    return Ok(c);
                }
            }
            Err(e) => return Err(e),
        }
    }

    Ok(0)
}

fn main() -> Result<(), Error> {
    let mut file = File::open("..\\..\\input.txt")?;
    let r_floor = solve_floor(&mut file)?;
    file.seek(std::io::SeekFrom::Start(0))?;
    let r_basement = solve_basement(&mut file)?;
    println!("Floor: {}", r_floor);
    println!("Basement: {}", r_basement);

    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;
    use std::io::{Error, Read};

    struct TestFloor(&'static str, i64);

    #[test]
    fn test_solve_floor() -> Result<(), Error> {
        let inputs = [
            TestFloor("", 0),
            TestFloor("(", 1),
            TestFloor("((", 2),
            TestFloor(")", -1),
            TestFloor("()", 0),
        ];

        for i in inputs {
            let r = solve_floor(i.0.as_bytes())?;
            assert_eq!(i.1, r);
        }
        Ok(())
    }
}
