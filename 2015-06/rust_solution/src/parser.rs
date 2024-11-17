use crate::range::{self, Point};
use std::error::Error;
use std::fmt::Display;
use std::{iter::Peekable, str::Chars};

#[derive(Debug, PartialEq)]
pub struct ParseError;

impl Display for ParseError {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "Parsing error")
    }
}

impl Error for ParseError {}

#[derive(Copy, Clone, Debug, PartialEq)]
pub enum Operation {
    TurnOn,
    TurnOff,
    Toggle,
}

#[derive(Copy, Clone, Debug, PartialEq)]
pub struct ParsedLine {
    pub op: Operation,
    pub from: range::Point,
    pub to: range::Point,
}

#[derive(Debug)]
pub struct Parser<'a> {
    start: usize,
    next_ind: usize,
    src: &'a str,
    chars: Peekable<Chars<'a>>,
}

impl<'a> Parser<'a> {
    pub fn new(input: &'a str) -> Self {
        Self {
            start: 0,
            next_ind: 0,
            src: input,
            chars: input.chars().peekable(),
        }
    }

    pub fn parse(&mut self) -> Result<ParsedLine, ParseError> {
        // Operation one or 2 words
        let op_token1 = self.next_token().ok_or(ParseError {})?;
        let op_token2 = self.next_token().ok_or(ParseError {})?;
        let op = Self::get_operation(op_token1, op_token2)?;

        // First point
        let p1_token1 = if op != Operation::Toggle {
            self.next_token().ok_or(ParseError {})?
        } else {
            op_token2
        };
        let p1_token2 = self.next_token().ok_or(ParseError {})?;
        let from = Self::get_point(p1_token1, p1_token2)?;

        // Skip "through"
        let _ = self.next_token().ok_or(ParseError {})?;

        //Second point
        let p2_token1 = self.next_token().ok_or(ParseError {})?;
        let p2_token2 = self.next_token().ok_or(ParseError {})?;
        let to = Self::get_point(p2_token1, p2_token2)?;

        Ok(ParsedLine { op, from, to })
    }

    fn get_operation(token1: &str, token2: &str) -> Result<Operation, ParseError> {
        match (token1, token2) {
            ("turn", "on") => Ok(Operation::TurnOn),
            ("turn", "off") => Ok(Operation::TurnOff),
            ("toggle", _) => Ok(Operation::Toggle),
            _ => Err(ParseError {}),
        }
    }

    fn get_point(token1: &str, token2: &str) -> Result<Point, ParseError> {
        let x: usize = token1.parse().map_err(|_| ParseError {})?;
        let y: usize = token2.parse().map_err(|_| ParseError {})?;
        Ok(Point(x, y))
    }

    fn skip_rubbish(&mut self) {
        while let Some(peeked) = self.chars.peek() {
            if !peeked.is_alphanumeric() {
                self.next_ind += 1;
                self.chars.next();
            } else {
                break;
            }
        }
    }

    fn next_token(&mut self) -> Option<&'a str> {
        self.skip_rubbish();
        self.start = self.next_ind;
        loop {
            self.next_ind += 1;
            let c = self.chars.by_ref().next();
            if let Some(c) = c {
                if !c.is_alphanumeric() {
                    break;
                }
            } else {
                break;
            }
        }

        if self.start < self.next_ind - 1 {
            Some(self.exctract_substring(self.start, self.next_ind - 1))
        } else {
            None
        }
    }

    fn exctract_substring(&self, start: usize, end: usize) -> &'a str {
        let (start_i, _) = self.src.char_indices().nth(start).unwrap_or((0, 0 as char));
        let (stop_i, _) = self
            .src
            .char_indices()
            .nth(end)
            .unwrap_or((self.src.len(), 0 as char));
        &self.src[start_i..stop_i]
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test_skip_rubbish() {
        let mut parser = Parser::new(" abc");
        parser.skip_rubbish();
        assert_eq!(1, parser.next_ind);
        assert_eq!('a', parser.chars.next().unwrap());
    }

    #[test]
    fn test_extract_substring() {
        let mut parser = Parser::new("aaa bbb");
        assert_eq!("aaa", parser.exctract_substring(0, 3));
        assert_eq!("bbb", parser.exctract_substring(4, 8));
    }

    #[test]
    fn test_next_token() {
        let mut empty_parser = Parser::new("");
        assert_eq!(None, empty_parser.next_token());
        assert_eq!(None, empty_parser.next_token());

        let mut tokens = Parser::new("ab cd");
        assert_eq!(Some("ab"), tokens.next_token());
        assert_eq!(Some("cd"), tokens.next_token());
        assert_eq!(None, tokens.next_token());

        let mut commas = Parser::new("ab,cd");
        assert_eq!(Some("ab"), commas.next_token());
        assert_eq!(Some("cd"), commas.next_token());
        assert_eq!(None, commas.next_token());
    }

    #[test]
    fn test_parsing_turn_on() {
        let mut parser = Parser::new("turn on 0,0 through 999,999");
        let expected = ParsedLine{
            op: Operation::TurnOn,
            from: Point(0, 0),
            to: Point(999, 999),
        };
        assert_eq!(Ok(expected), parser.parse());
    }

    
    #[test]
    fn test_parsing_turn_off() {
        let mut parser = Parser::new("turn off 100,0 through 999,0");
        let expected = ParsedLine{
            op: Operation::TurnOff,
            from: Point(100, 0),
            to: Point(999, 0),
        };
        assert_eq!(Ok(expected), parser.parse());
    }

    #[test]
    fn test_parsing_toggle() {
        let mut parser = Parser::new("toggle 0,0 through 999,999");
        let expected = ParsedLine{
            op: Operation::Toggle,
            from: Point(0, 0),
            to: Point(999, 999),
        };
        assert_eq!(Ok(expected), parser.parse());
    }

    #[test]
    fn nth() {
        let s = "привет";
        for (i, c) in s.char_indices() {
            println!("{}={}", i, c);
        }
        let a = &s[0..4];
        println!("{}", a);
    }

}
