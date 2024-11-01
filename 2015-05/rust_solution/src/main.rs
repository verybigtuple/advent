use std::collections::HashMap;
use std::fs::File;
use std::io::{self, BufRead};

fn is_nice1(input: &str) -> bool {
    let mut prev: Option<char> = None;
    let mut vowels = 0;
    let mut pairs = 0;
    for c in input.chars() {
        match c {
            'a' | 'e' | 'i' | 'o' | 'u' => vowels += 1,
            'b' | 'd' | 'q' | 'y' => {
                if let Some(prev_char) = prev {
                    if matches!(
                        (prev_char, c),
                        ('a', 'b') | ('c', 'd') | ('p', 'q') | ('x', 'y')
                    ) {
                        return false;
                    }
                }
            }
            _ => {}
        }
        if Some(c) == prev {
            pairs += 1;
        }

        prev = Some(c);
    }

    vowels >= 3 && pairs >= 1
}

#[derive(Clone, Copy, Debug)]
struct Buf {
    b: [Option<char>; 2],
}

impl Buf {
    fn new() -> Buf {
        Buf { b: [None, None] }
    }

    fn push(&mut self, c: char) {
        (self.b[0], self.b[1]) = (Some(c), self.b[0])
    }
}

#[derive(Clone, Copy, Hash, Eq, PartialEq, Debug)]
struct Pair(char, char);

fn is_nice2(input: &str) -> bool {
    let mut buf = Buf::new();
    let mut has_pair = false;
    let mut has_repeat = false;
    let mut pairs: HashMap<Pair, usize> = HashMap::new();
    for (i, c) in input.chars().enumerate() {
        if !has_pair {
            if let Some(p0) = buf.b[0] {
                let current_pair = Pair(p0, c);
                match pairs.get(&current_pair) {
                    Some(ind) => {
                        if i - ind > 1 {
                            has_pair = true;
                        }
                    }
                    None => {
                        pairs.insert(current_pair, i);
                    }
                }
            }
        }

        if !has_repeat {
            if let Some(p1) = buf.b[1] {
                has_repeat = p1 == c;
            }
        }

        if has_pair && has_repeat {
            return true;
        }

        buf.push(c);
    }
    false
}

fn main() -> io::Result<()> {
    let file = File::open("input.txt")?;
    let reader = io::BufReader::new(file);
    let mut nice1 = 0;
    let mut nice2 = 0;
    for line in reader.lines() {
        let line = line?;
        if is_nice1(&line) {
            nice1 += 1;
        }
        if is_nice2(&line) {
            nice2 += 1;
        }
    }
    println!("Nice1 {}", nice1);
    println!("Nice2 {}", nice2);

    Ok(())
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_is_nice1() {
        assert!(is_nice1("aaa"));
        assert!(is_nice1("ugknbfddgicrmopn"));
        assert!(!is_nice1("jchzalrnumimnmhp"));
        assert!(!is_nice1("haegwjzuvuyypxyu"));
        assert!(!is_nice1("dvszwmarrgswjxmb"));
    }

    #[test]
    fn test_is_nice2() {
        assert!(is_nice2("qjhvhtzxzqqjkmpb"));
        assert!(is_nice2("xxyxx"));
        assert!(is_nice2("xxxx"));
        assert!(!is_nice2("xxx"));
        assert!(!is_nice2("uurcxstgmygtbstg"));
        assert!(!is_nice2("uurcxstgmygtbstg"));
    }
}
