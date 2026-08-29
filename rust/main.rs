use std::sync::mpsc::{Receiver, Sender, channel};
use std::thread;

fn printer(start_number: i32, receiver: Receiver<()>, sender: Sender<()>) {
    for i in (start_number..=100).step_by(2) {
        receiver.recv().unwrap();
        println!("{i}");
        if i != 100 {
            sender.send(()).unwrap();
        }
    }
}

fn main() {
    let (odd_tx, odd_rx) = channel();
    let (even_tx, even_rx) = channel();
    let even_tx_1 = even_tx.clone();
    
    let odd_handle = thread::spawn(move || printer(1, odd_rx, even_tx_1));

    let even_handle = thread::spawn(move || printer(0, even_rx, odd_tx));

    even_tx.send(()).unwrap();

    odd_handle.join().unwrap();
    even_handle.join().unwrap();
}
