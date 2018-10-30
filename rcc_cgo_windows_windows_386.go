package main

/*
#cgo CFLAGS: -pipe -fno-keep-inline-dllexport -O2 -Wall -W -Wextra -DUNICODE -D_UNICODE -DWIN32 -DQT_NEEDS_QMAIN -DQT_NO_DEBUG -DQT_GUI_LIB -DQT_CORE_LIB
#cgo CXXFLAGS: -pipe -fno-keep-inline-dllexport -O2 -std=gnu++11 -Wall -W -Wextra -fexceptions -mthreads -DUNICODE -D_UNICODE -DWIN32 -DQT_NEEDS_QMAIN -DQT_NO_DEBUG -DQT_GUI_LIB -DQT_CORE_LIB
#cgo CXXFLAGS: -I../../src -I. -I/usr/lib/mxe/usr/i686-w64-mingw32.static/qt5/include -I/usr/lib/mxe/usr/i686-w64-mingw32.static/qt5/include/QtGui -I/usr/lib/mxe/usr/i686-w64-mingw32.static/qt5/include/QtCore -Irelease -I/usr/lib/mxe/usr/i686-w64-mingw32.static/qt5/mkspecs/win32-g++
#cgo LDFLAGS:        -Wl,-s -Wl,-subsystem,windows -mthreads
#cgo LDFLAGS:        -lmingw32 -L/usr/lib/mxe/usr/i686-w64-mingw32.static/qt5/lib /usr/lib/mxe/usr/i686-w64-mingw32.static/qt5/lib/libqtmain.a -L/usr/lib/mxe/usr/i686-w64-mingw32.static/qt5/plugins/platforms /usr/lib/mxe/usr/i686-w64-mingw32.static/qt5/plugins/platforms/libqwindows.a -ldwmapi -lwinspool /usr/lib/mxe/usr/i686-w64-mingw32.static/qt5/lib/libQt5EventDispatcherSupport.a -L/usr/lib/mxe/usr/i686-w64-mingw32.static/lib /usr/lib/mxe/usr/i686-w64-mingw32.static/qt5/lib/libQt5FontDatabaseSupport.a /usr/lib/mxe/usr/i686-w64-mingw32.static/qt5/lib/libQt5ThemeSupport.a /usr/lib/mxe/usr/i686-w64-mingw32.static/qt5/lib/libQt5AccessibilitySupport.a /usr/lib/mxe/usr/i686-w64-mingw32.static/qt5/lib/libQt5WindowsUIAutomationSupport.a -L/usr/lib/mxe/usr/i686-w64-mingw32.static/qt5/plugins/imageformats /usr/lib/mxe/usr/i686-w64-mingw32.static/qt5/plugins/imageformats/libqgif.a /usr/lib/mxe/usr/i686-w64-mingw32.static/qt5/plugins/imageformats/libqicns.a /usr/lib/mxe/usr/i686-w64-mingw32.static/qt5/plugins/imageformats/libqico.a /usr/lib/mxe/usr/i686-w64-mingw32.static/qt5/plugins/imageformats/libqjp2.a -ljasper /usr/lib/mxe/usr/i686-w64-mingw32.static/qt5/plugins/imageformats/libqjpeg.a /usr/lib/mxe/usr/i686-w64-mingw32.static/qt5/plugins/imageformats/libqmng.a -lmng -llcms2 -lpthread /usr/lib/mxe/usr/i686-w64-mingw32.static/qt5/plugins/imageformats/libqtga.a /usr/lib/mxe/usr/i686-w64-mingw32.static/qt5/plugins/imageformats/libqtiff.a -ltiff -llzma -ljpeg /usr/lib/mxe/usr/i686-w64-mingw32.static/qt5/plugins/imageformats/libqwbmp.a /usr/lib/mxe/usr/i686-w64-mingw32.static/qt5/plugins/imageformats/libqwebp.a /usr/lib/mxe/usr/i686-w64-mingw32.static/qt5/lib/libQt5Gui.a -lcomdlg32 -loleaut32 -limm32 -lopengl32 -lharfbuzz -lcairo -lgobject-2.0 -lfontconfig -lfreetype -lm -lusp10 -lmsimg32 -lgdi32 -lpixman-1 -lffi -lexpat -lbz2 -lpng16 -lharfbuzz_too -lfreetype_too -lglib-2.0 -lshlwapi -lpcre -lintl -liconv /usr/lib/mxe/usr/i686-w64-mingw32.static/qt5/lib/libQt5Core.a -lmpr -lnetapi32 -luserenv -lversion -lws2_32 -lkernel32 -luser32 -lshell32 -luuid -lole32 -ladvapi32 -lwinmm -lz -lpcre2-16
#cgo LDFLAGS: -Wl,--allow-multiple-definition
#cgo CFLAGS: -Wno-unused-parameter -Wno-unused-variable -Wno-return-type
#cgo CXXFLAGS: -Wno-unused-parameter -Wno-unused-variable -Wno-return-type
*/
import "C"