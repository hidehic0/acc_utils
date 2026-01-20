{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    gomod2nix = {
      url = "github:nix-community/gomod2nix";
      inputs.nixpkgs.follows = "nixpkgs";
      inputs.flake-utils.follows = "flake-utils";
    };
  };

  outputs =
    {
      nixpkgs,
      flake-utils,
      gomod2nix,
      ...
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs {
          inherit system;
          overlays = [
            gomod2nix.overlays.default
          ];
        };

        meta = {
          description = "A tool with enhanced functionality using files created by atcoder-cli";
          homepage = "https://github.com/hidehic0/acc_utils";
          license = pkgs.lib.licenses.unlicense;
        };
      in
      {
        packages = {
          default = pkgs.buildGoApplication {
            name = "acc_utils";
            meta = meta;
            src = ./.;
            modules = ./gomod2nix.toml;
            postInstall = ''
              mkdir -p $out/share/zsh/site-functions
              $out/bin/acc_utils completion zsh > $out/share/zsh/site-functions/_acc_utils
            '';
          };
        };
        devShells.default = pkgs.mkShell {
          packages = [
            pkgs.goreleaser
            pkgs.go
            gomod2nix.packages.${system}.default
          ];
        };
      }
    );
}
