with import <nixpkgs> {};
mkShell {
  nativeBuildInputs = [
    bashInteractive
    go
  ];
}
