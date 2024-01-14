# ransomware-demo

## Description
- This repo features a demo ransomware implementation written in Go.
- It consists of two files: `enc.go` (used for data encryption) and `dec.go` (used for data decryption).
- Program relies on AES-128 decryption in GCM mode.

## Usage
```
go run enc.go
```
This will run an encryption program that would encrypt all of the data in the provided path
```
go run dec.go
```
This will run a decryption program that would decrypt all of the previously decrypted data in the provided path

## Good to know
- Since it is made as a demo, it also has a `fake_drive` directory inside a project that will be uses as a test drive/directory for running the ransomware on. Ofcourse, you can change it to any other path for your usecase (keep in mind that if using for something like `C:\\` you will most likely need to worry about priviledges.
- Also for demo purposes, it is made so that only pdf files will be encrypted. You can change `enc.go` so that it would encrypt any other extension, rewrite it so that it takes multiple extensions or just ommit the extension check. However, in `dec.go` the extension check is important since it is only ment to run on files that were previously encrypted

## Example Output
```
Encrypting 'fake_drive' with AES-128...
[!] Encrypting 'sample_1.pdf'... DONE
[!] Encrypting 'sample_2.pdf'... DONE
------------------------------------------
Encryption finished successfully!
Duration: 519.968µs
```

```
Decrypting 'fake_drive' with AES-128...
[!] Decrypting 'sample_1.pdf.enc'... DONE
[!] Decrypting 'sample_2.pdf.enc'... DONE
------------------------------------------
Decryption finished successfully!
Duration: 382.534µs
```
