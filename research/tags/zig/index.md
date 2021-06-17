+++
title = "Zig"
tags = [ "proposal" ]
date = 2020-10-05
+++

## Download

<https://github.com/ziglang/zig>

## Static OpenSSL Windows

It should be possible to create a Static Windows build with OpenSSL. It seems
currently only Shared (DLL) builds are possible. Static Windows builds of
OpenSSL are available here:

- <https://github.com/marler8997/ziget/blob/3daa9be4/build.zig#L66-L74>
- <https://packages.msys2.org/package/mingw-w64-x86_64-openssl>

To anyone interested, I was able to get this working like this:

~~~diff
diff --git a/build.zig b/build.zig
index 4d06a09..504808e 100644
--- a/build.zig
+++ b/build.zig
@@ -29,9 +29,11 @@ pub fn build(b: *Builder) !void {
         exe.addPackage(Pkg { .name = "ssl", .path = "openssl/ssl.zig" });
         exe.linkSystemLibrary("c");
         if (std.builtin.os.tag == .windows) {
-            exe.linkSystemLibrary("libcrypto");
-            exe.linkSystemLibrary("libssl");
-            try setupOpensslWindows(b, exe);
+            exe.addObjectFile("D:/Desktop/mingw64/lib/libssl.a");
+            exe.addObjectFile("D:/Desktop/mingw64/lib/libcrypto.a");
+            exe.addIncludeDir("D:/Desktop/mingw64/include");
+            exe.addObjectFile("C:/msys2/mingw64/x86_64-w64-mingw32/lib/libmincore.a");
+            exe.addObjectFile("C:/msys2/mingw64/x86_64-w64-mingw32/lib/libws2_32.a");
~~~

~~~
zig.exe build -Dopenssl=true -Dtarget=native-native-gnu
ziget.exe --stdout https://speedtest.lax.hivelocity.net
~~~

this is a bad fix as it clobbers the existing option of Shared linking.
