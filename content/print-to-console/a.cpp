#include <codecvt>
#include <iostream>

int main() {
   std::ios_base::sync_with_stdio(false);
   std::locale utf8(std::locale(), new std::codecvt_utf8_utf16<wchar_t>);
   std::wcout.imbue(utf8);
   std::wcout << L"📗📕" << std::endl;
}
