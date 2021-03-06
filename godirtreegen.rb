# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Godirtreegen < Formula
  desc ""
  homepage "https://github.com/rexsimiloluwah/godirtreegen"
  version "0.1.7"

  on_macos do
    url "https://github.com/rexsimiloluwah/godirtreegen/releases/download/v0.1.7/godirtreegen_0.1.7_darwin_all.tar.gz"
    sha256 "a49db9b01a8e876ad2e1c0383ca85e71a5ac55015da18202a97eb3a9c901c080"

    def install
      bin.install "godirtreegen"
    end
  end

  on_linux do
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/rexsimiloluwah/godirtreegen/releases/download/v0.1.7/godirtreegen_0.1.7_linux_arm64.tar.gz"
      sha256 "9d4103431aff59b0d6fb78b3cae11a13f4e435d612dc7aa51ab18386f20ee92e"

      def install
        bin.install "godirtreegen"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/rexsimiloluwah/godirtreegen/releases/download/v0.1.7/godirtreegen_0.1.7_linux_amd64.tar.gz"
      sha256 "27627d5a174fbdb28e8a24995485cf79f519f7ec8647dcf3b8b3341b77803d24"

      def install
        bin.install "godirtreegen"
      end
    end
    if Hardware::CPU.arm? && !Hardware::CPU.is_64_bit?
      url "https://github.com/rexsimiloluwah/godirtreegen/releases/download/v0.1.7/godirtreegen_0.1.7_linux_armv6.tar.gz"
      sha256 "feab49defe21c4e1bc5676e63998cbf57459871df488b50557eac77d507ca52e"

      def install
        bin.install "godirtreegen"
      end
    end
  end
end
