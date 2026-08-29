use std::sync::mpsc;
use std::thread;

fn main() {
    let (odd_tx, odd_rx) = mpsc::channel();
    let (even_tx, even_rx) = mpsc::channel();
    let even_tx_1 = even_tx.clone();

    let odd_handle = thread::spawn(move || {
        for i in (1..=100).step_by(2) {
            odd_rx.recv().unwrap();
            println!("{i}");
            even_tx_1.send(()).unwrap();
        }
    });

    let even_handle = thread::spawn(move || {
        for i in (0..=100).step_by(2) {
            even_rx.recv().unwrap();
            println!("{i}");
            if i != 100 {
                odd_tx.send(()).unwrap();
            }
        }
    });

    even_tx.send(()).unwrap();

    odd_handle.join().unwrap();
    even_handle.join().unwrap();
}