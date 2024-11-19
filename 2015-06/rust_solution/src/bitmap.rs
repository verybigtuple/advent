const ITEM_BITS: usize = 8;

#[derive(Debug)]
pub struct BitMap {
    len: usize,
    vec: Vec<u8>,
    count: usize,
}

impl BitMap {
    pub fn new(bit_len: usize) -> BitMap {
        let vec_len = Self::calc_internal_size(bit_len);
        return BitMap {
            len: bit_len,
            vec: vec![0; vec_len],
            count: 0,
        };
    }

    fn calc_internal_size(bit_len: usize) -> usize {
        (bit_len - 1) / ITEM_BITS + 1
    }

    fn get_item_index(bit_n: usize) -> usize {
        bit_n / ITEM_BITS
    }

    fn shift(bit_n: usize) -> u8 {
        1 << (bit_n % ITEM_BITS)
    }

    fn check_size(&self, bit_n: usize) {
        if bit_n > self.len {
            panic!("Out of bounds");
        }
    }

    pub fn count(&self) -> usize {
        self.count
    }

    pub fn set_bit(&mut self, bit_n: usize) {
        self.check_size(bit_n);
        let i = Self::get_item_index(bit_n);
        if !self.get_bit(bit_n) {
            self.vec[i] |= Self::shift(bit_n);
            self.count += 1;
        }
    }

    pub fn reset_bit(&mut self, bit_n: usize) {
        self.check_size(bit_n);
        let i = Self::get_item_index(bit_n);
        if self.get_bit(bit_n) {
            self.vec[i] &= !Self::shift(bit_n);
            self.count -= 1;
        }
    }

    pub fn get_bit(&self, bit_n: usize) -> bool {
        self.check_size(bit_n);
        let i = Self::get_item_index(bit_n);
        (self.vec[i] & Self::shift(bit_n)) != 0
    }

    pub fn toggle_bit(&mut self, bit_n: usize) {
        self.check_size(bit_n);
        let i = Self::get_item_index(bit_n);
        // No XOR as I want to count bits
        if self.get_bit(bit_n) {
            self.vec[i] &= !Self::shift(bit_n);
            self.count -= 1;
        } else {
            self.vec[i] |= Self::shift(bit_n);
            self.count += 1;
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_calc_size() {
        assert_eq!(1, BitMap::calc_internal_size(1), "1 flag");
        assert_eq!(1, BitMap::calc_internal_size(8), "8 flags");
        assert_eq!(2, BitMap::calc_internal_size(9));
        assert_eq!(2, BitMap::calc_internal_size(15));
        assert_eq!(2, BitMap::calc_internal_size(16));
        assert_eq!(3, BitMap::calc_internal_size(17));
    }

    #[test]
    fn test_get_item_index() {
        assert_eq!(0, BitMap::get_item_index(0));
        assert_eq!(0, BitMap::get_item_index(1));
        assert_eq!(0, BitMap::get_item_index(7));
        assert_eq!(1, BitMap::get_item_index(8));
    }

    #[test]
    fn test_zeros() {
        let bm = BitMap::new(8);
        for i in 0..8 {
            assert!(!bm.get_bit(i), "i={}", i)
        }
    }

    #[test]
    fn test_set() {
        let mut bm = BitMap::new(9);
        bm.set_bit(0);
        assert_eq!(0b0000_0001, bm.vec[0]);
        bm.set_bit(7);
        assert_eq!(0b1000_0001, bm.vec[0]);
        bm.set_bit(8);
        assert_eq!(0b0000_0001, bm.vec[1]);
    }

    #[test]
    fn test_get() {
        let mut bm = BitMap::new(8);
        bm.set_bit(0);
        bm.set_bit(7);
        assert!(bm.get_bit(0));
        assert!(bm.get_bit(7));
        assert!(!bm.get_bit(6));
    }

    #[test]
    fn test_reset() {
        let mut bm = BitMap::new(8);
        bm.set_bit(7);
        assert!(bm.get_bit(7));
        bm.reset_bit(7);
        assert!(!bm.get_bit(7));
    }

    #[test]
    fn test_toggle() {
        let mut bm = BitMap::new(8);
        bm.toggle_bit(7);
        assert!(bm.get_bit(7));
        assert!(!bm.get_bit(0));
        bm.toggle_bit(7);
        assert!(!bm.get_bit(7));
        assert!(!bm.get_bit(0));
    }

    #[test]
    fn test_count() {
        let mut bm = BitMap::new(8);
        assert_eq!(0, bm.count());
        bm.set_bit(0);
        assert_eq!(1, bm.count());
        bm.set_bit(0);
        assert_eq!(1, bm.count());
        bm.set_bit(1);
        assert_eq!(2, bm.count());
        bm.reset_bit(1);
        assert_eq!(1, bm.count());
        bm.toggle_bit(0);
        assert_eq!(0, bm.count());
    }
}
