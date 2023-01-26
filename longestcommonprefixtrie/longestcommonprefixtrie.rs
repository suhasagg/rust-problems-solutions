impl Solution {
    pub fn longest_common_prefix(strs: Vec<String>) -> String {
    let mut trie = Trie::default();
        for word in strs.iter() {
        trie.insert(word);
    }
     return trie.lcp();
    }
}

use std::collections::HashMap;

#[derive(Default)]
struct Trie {
    children: HashMap<char, Box<Trie>>,
    is_leaf: bool,
}

impl Trie {
    fn insert(&mut self, word: &str) {
        let mut curr = self;
        for c in word.chars() {
            curr = curr.children.entry(c).or_insert(Box::new(Trie::default()));
        }
        curr.is_leaf = true;
    }

    fn lcp(&self) -> String {
        let mut curr = self;
        let mut prefix = String::new();
        while !curr.is_leaf && curr.children.len() == 1 {
            for (c, child) in curr.children.iter() {
                prefix.push(*c);
                curr = child;
            }
        }
        prefix
    }
}
