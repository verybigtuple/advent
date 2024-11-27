#[derive(Clone, Copy, Debug, PartialEq, PartialOrd)]
pub struct Point(pub usize, pub usize);

#[derive(Clone, Copy, Debug)]
pub struct PointRange {
    from: Point,
    to: Point,
    max: Point,
    x: usize,
    y: usize,
}

impl PointRange {
    pub fn new(from: Point, to: Point, max: Point) -> PointRange {
        if from > to {
            panic!("To point should be more than From point");
        }

        if from > max || to > max {
            panic!("points cannot be more than max");
        }

        PointRange {
            from,
            to,
            max,
            x: from.0,
            y: from.1,
        }
    }
}

impl Iterator for PointRange {
    type Item = usize;

    fn next(&mut self) -> Option<Self::Item> {
        if self.y > self.to.1 {
            return None;
        }

        let current = self.x + self.y * (self.max.0 + 1);

        if self.x < self.to.0 {
            self.x += 1;
        } else {
            self.x = self.from.0;
            self.y += 1;
        }

        Some(current)
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_empty_range() {
        let start = Point(0, 0);
        let mut r = PointRange::new(start, start, start);
        assert_eq!(Some(0), r.next());
        assert_eq!(None, r.next());
    }

    #[test]
    fn test_one_point() {
        let mut r = PointRange::new(Point(1, 1), Point(1, 1), Point(2, 2));
        assert_eq!(Some(4), r.next());
        assert_eq!(None, r.next());
    }

    #[test]
    fn test_range() {
        let mut r = PointRange::new(Point(1, 1), Point(2, 2), Point(2, 2));
        for i in (4..=8).map(|x| x as usize) {
            assert_eq!(Some(i), r.next());
        }
        assert_eq!(None, r.next());
    }

    #[test]
    fn test_count_big_range() {
        let r = PointRange::new(Point(0, 0), Point(999, 999), Point(999, 999));
        assert_eq!(1_000_000, r.count());
    }

    #[test]
    fn test_count_null_range() {
        let r = PointRange::new(Point(499, 499), Point(500, 500), Point(999, 999));
        assert_eq!(4, r.count());
    }
}
