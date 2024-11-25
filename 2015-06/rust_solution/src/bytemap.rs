#[derive(Debug)]
pub struct ByteMap {
    vec: Vec<u8>,
    count: usize,
}

impl ByteMap {
    pub fn new(len: usize) -> ByteMap {
        return ByteMap {
            vec: vec![0; len],
            count: 0,
        };
    }

    pub fn count(&self) -> usize {
        self.count
    }

    pub fn inc_byte(&mut self, byte_n: usize) {
        self.inc_byte_by(byte_n, 1);
    }

    pub fn inc_byte_by(&mut self, byte_n: usize, by: u8) {
        self.vec[byte_n] += by;
        self.count += by as usize;
    }

    pub fn dec_byte(&mut self, byte_n: usize) {
        if self.vec[byte_n] > 0 {
            self.vec[byte_n] -= 1;
            self.count -= 1;
        }
    }
}
