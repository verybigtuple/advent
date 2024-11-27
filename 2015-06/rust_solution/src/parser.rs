use crate::range::Point;
use std::error::Error;
use std::fmt::Display;
use std::iter::Peekable;
use std::str::CharIndices;

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
    pub from: Point,
    pub to: Point,
}

#[derive(Debug)]
pub struct Parser<'a> {
    src: &'a str,
    chars: Peekable<CharIndices<'a>>,
}

impl<'a> Parser<'a> {
    pub fn new(input: &'a str) -> Self {
        Self {
            src: input,
            chars: input.char_indices().peekable(),
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

    fn next_token(&mut self) -> Option<&'a str> {
        // Skip non-alphanumeric characters
        let start = self
            .chars
            .by_ref()
            .find(|(_, c)| c.is_alphanumeric())
            .map(|(ind, _)| ind)?; // if no tokens here we return

        // Find the end of the token (inclusive)
        let end = self
            .chars
            .by_ref()
            .find(|(_, c)| !c.is_alphanumeric())
            .map(|(ind, _)| ind)
            .unwrap_or(self.src.len());

        // Return the token slice using the valid UTF-8 indices
        Some(&self.src[start..end])
    }
}

#[cfg(test)]
mod tests {
    use super::*;

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
    fn test_utf_tokens() {
        let mut tokens = Parser::new("привет, мир");
        assert_eq!(Some("привет"), tokens.next_token());
        assert_eq!(Some("мир"), tokens.next_token());
        assert_eq!(None, tokens.next_token());
    }

    #[test]
    fn test_parsing_turn_on() {
        let mut parser = Parser::new("turn on 0,0 through 999,999");
        let expected = ParsedLine {
            op: Operation::TurnOn,
            from: Point(0, 0),
            to: Point(999, 999),
        };
        assert_eq!(Ok(expected), parser.parse());
    }

    #[test]
    fn test_parsing_turn_off() {
        let mut parser = Parser::new("turn off 100,0 through 999,0");
        let expected = ParsedLine {
            op: Operation::TurnOff,
            from: Point(100, 0),
            to: Point(999, 0),
        };
        assert_eq!(Ok(expected), parser.parse());
    }

    #[test]
    fn test_parsing_toggle() {
        let mut parser = Parser::new("toggle 0,0 through 999,999");
        let expected = ParsedLine {
            op: Operation::Toggle,
            from: Point(0, 0),
            to: Point(999, 999),
        };
        assert_eq!(Ok(expected), parser.parse());
    }
}
