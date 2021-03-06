package mimemagician

import (
	"testing"
)

func TestMatchGlob(t *testing.T) {
	tests := []struct {
		filename string
		want     string
	}{
		{"filename.extension", "application/octet-stream"},
		{"meson_options.txt", "text/x-meson"},
		{"text.txt", "text/plain"},
		{"makefile.conf", "text/x-makefile"},
		{"video.webm", "video/webm"},
		{"audio.mka", "audio/x-matroska"},
		{"sourcecode.C", "text/x-c++src"},
		{"358.vdr", "video/mpeg"},
		{"animation.anim4", "video/x-anim"},
		{"file.animj", "video/x-anim"},
		{"file.anim0", "application/octet-stream"},
		{"trashfile~", "application/x-trash"},
		{"settings.conf.bak", "application/x-trash"},
		{"doom.wad", "application/x-doom-wad"},
		{"2001_compression_overview.djvu", "image/vnd.djvu"},
		{"32x-rom.32x", "application/x-genesis-32x-rom"},
		{"4jsno.669", "audio/x-mod"},
		{"560051.xml", "application/xml"},
		{"adf-test.adf", "application/x-amiga-disk-format"},
		{"aero_alt.cur", "image/x-win-bitmap"},
		{"all_w.m3u8", "application/vnd.apple.mpegurl"},
		{"Anaphraseus-1.21-beta.oxt", "application/vnd.openofficeorg.extension"},
		{"ancp.pcap", "application/vnd.tcpdump.pcap"},
		{"androide.k7", "application/x-thomson-cassette"},
		{"aportis.pdb", "application/vnd.palm"},
		{"archive.7z", "application/x-7z-compressed"},
		{"archive.lrz", "application/x-lrzip"},
		{"archive.tar", "application/x-tar"},
		{"ascii.stl", "model/stl"},
		{"atari-2600-test.A26", "application/x-atari-2600-rom"},
		{"atari-7800-test.A78", "application/x-atari-7800-rom"},
		{"atari-lynx-chips-challenge.lnx", "application/x-atari-lynx-rom"},
		{"attachment.tif", "image/tiff"},
		{"balloon.j2c", "image/x-jp2-codestream"},
		{"balloon.jp2", "image/jp2"},
		{"balloon.jpf", "image/jpx"},
		{"balloon.jpm", "image/jpm"},
		{"balloon.mj2", "video/mj2"},
		{"bathead.sk", "image/x-skencil"},
		{"bbc.ram", "application/ram"},
		{"bibtex.bib", "text/x-bibtex"},
		{"binary.stl", "model/stl"},
		{"blitz.m7", "application/x-thomson-cartridge-memo7"},
		{"bluerect.mdi", "image/vnd.ms-modi"},
		{"bluish.icc", "application/vnd.iccprofile"},
		{"break.mtm", "audio/x-mod"},
		{"bug106330.iso", "application/x-cd-image"},
		{"bug-30656-xchat.conf", "application/octet-stream"},
		{"bug39126-broken.ppm", "image/x-portable-pixmap"},
		{"bug39126-working.ppm", "image/x-portable-pixmap"},
		{"build.gradle", "text/x-gradle"},
		{"ccfilm.axv", "video/annodex"},
		{"classiq1.hfe", "application/x-hfe-floppy-image"},
		{"colormapped.tga", "image/x-tga"},
		{"combined.karbon", "application/x-karbon"},
		{"comics.cb7", "application/x-cb7"},
		{"comics.cbt", "application/x-cbt"},
		{"COPYING.asc", "text/plain"},
		{"copying.cab", "application/vnd.ms-cab-compressed"},
		{"COPYING-clearsign.asc", "text/plain"},
		{"COPYING-encrypted.asc", "text/plain"},
		{"core", "application/x-core"},
		{"Core", "application/octet-stream"},
		{"ct_faac-adts.aac", "audio/aac"},
		{"cube.igs", "model/iges"},
		{"cube.wrl", "model/vrml"},
		{"cyborg.med", "audio/x-mod"},
		{"dbus-comment.service", "text/x-dbus-service"},
		{"dbus.service", "text/x-dbus-service"},
		{"debian-goodies_0.63_all.deb", "application/vnd.debian.binary-package"},
		{"dia.shape", "application/x-dia-shape"},
		{"disk.img", "application/x-raw-disk-image"},
		{"disk.img.xz", "application/x-raw-disk-image-xz-compressed"},
		{"disk.raw-disk-image", "application/x-raw-disk-image"},
		{"disk.raw-disk-image.xz", "application/x-raw-disk-image-xz-compressed"},
		{"dns.cap", "application/vnd.tcpdump.pcap"},
		{"editcopy.png", "image/png"},
		{"Elephants_Dream-360p-Stereo.webm", "video/webm"},
		{"Empty.chrt", "application/x-kchart"},
		{"en_US.zip.meta4", "application/metalink4+xml"},
		{"esm.mjs", "application/javascript"},
		{"evolution.eml", "message/rfc822"},
		{"example_42_all.snap", "application/vnd.snap"},
		{"example.heic", "image/heif"},
		{"example.heif", "image/heif"},
		{"feed2", "application/octet-stream"},
		{"feed.atom", "application/atom+xml"},
		{"feed.rss", "application/rss+xml"},
		{"feeds.opml", "text/x-opml+xml"},
		{"foo-0.1-1.fc18.src.rpm", "application/x-source-rpm"},
		{"foo.doc", "application/msword"},
		{"fuji.themepack", "application/x-windows-themepack"},
		{"game-boy-color-test.gbc", "application/x-gameboy-color-rom"},
		{"game-boy-test.gb", "application/x-gameboy-rom"},
		{"game-gear-test.gg", "application/x-gamegear-rom"},
		{"GammaChart.exr", "image/x-exr"},
		{"gedit.flatpakref", "application/vnd.flatpak.ref"},
		{"genesis1.bin", "application/x-saturn-rom"},
		{"genesis2.bin", "application/x-saturn-rom"},
		{"gnome.flatpakrepo", "application/vnd.flatpak.repo"},
		{"good-1-delta-lzma2.tiff.xz", "application/x-xz"},
		{"googleearth.kml", "application/vnd.google-earth.kml+xml"},
		{"gtk-builder.ui", "application/x-designer"},
		{"hbo-playlist.qtl", "application/x-quicktime-media-link"},
		{"hello.flatpak", "application/vnd.flatpak"},
		{"hello.pack", "application/x-java-pack200"},
		{"helloworld.groovy", "text/x-groovy"},
		{"helloworld.java", "text/x-java"},
		{"helloworld.xpi", "application/x-xpinstall"},
		{"hello.xdgapp", "application/vnd.flatpak"},
		{"hereyes_remake.mo3", "audio/x-mo3"},
		{"html4.css", "text/css"},
		{"html5.css", "text/css"},
		{"image.sqsh", "application/vnd.squashfs"},
		{"img_5304.jpg", "image/jpeg"},
		{"internet.ez", "application/andrew-inset"},
		{"isdir.m", "application/vnd.wolfram.mathematica.package"},
		{"ISOcyr1.ent", "application/xml-external-parsed-entity"},
		{"iso-file.iso", "application/x-cd-image"},
		{"IWAD.WAD", "application/x-doom-wad"},
		{"javascript-without-extension", "application/octet-stream"},
		{"jc-win.ani", "application/x-navi-animation"},
		{"json_array.json", "application/json"},
		{"json-ld-full-iri.jsonld", "application/ld+json"},
		{"json_object.json", "application/json"},
		{"layersupdatesignals.flw", "application/x-kivio"},
		{"Leafpad-0.8.17-x86_64.AppImage", "application/vnd.appimage"},
		{"Leafpad-0.8.18.1.glibc2.4-x86_64.AppImage", "application/vnd.appimage"},
		{"libcompat.a", "application/x-archive"},
		{"libcompat.ar", "application/x-archive"},
		{"LiberationSans-Regular.ttf", "font/ttf"},
		{"LiberationSans-Regular.woff", "font/woff"},
		{"linguist.ts", "text/vnd.qt.linguist"},
		{"list", "application/octet-stream"},
		{"live-streaming.m3u", "application/vnd.apple.mpegurl"},
		{"lucid-tab-bg.xcf", "image/x-xcf"},
		{"m64p_test_rom.n64", "application/x-n64-rom"},
		{"m64p_test_rom.v64", "application/x-n64-rom"},
		{"m64p_test_rom.z64", "application/x-n64-rom"},
		{"Makefile", "text/x-makefile"},
		{"Makefile.gnu", "text/x-makefile"},
		{"markdown.md", "text/markdown"},
		{"mega-drive-rom.gen", "application/x-genesis-rom"},
		{"menu.ini", "application/octet-stream"},
		{"meson.build", "text/x-meson"},
		{"meson_options.txt", "text/x-meson"},
		{"Metroid_japan.fds", "application/x-fds-disk"},
		{"msg0001.gsm", "audio/x-gsm"},
		{"msx2-metal-gear.msx", "application/x-msx-rom"},
		{"msx-penguin-adventure.msx", "application/x-msx-rom"},
		{"my-data.json-patch", "application/json-patch+json"},
		{"mypaint.ora", "image/openraster"},
		{"mysum.m", "application/vnd.wolfram.mathematica.package"},
		{"neo-geo-pocket-color-test.ngc", "application/x-neo-geo-pocket-color-rom"},
		{"neo-geo-pocket-test.ngp", "application/x-neo-geo-pocket-rom"},
		{"newtonme.pict", "image/x-pict"},
		{"nrl.trig", "application/trig"},
		{"ocf10-20060911.epub", "application/epub+zip"},
		{"office.doc", "application/msword"},
		{"one-file.tnef", "application/vnd.ms-tnef"},
		{"ooo25876-2.pct", "image/x-pict"},
		{"ooo-6.0.doc", "application/msword"},
		{"ooo-95.doc", "application/msword"},
		{"ooo.doc", "application/msword"},
		{"ooo.rtf", "application/rtf"},
		{"ooo.sdw", "application/vnd.stardivision.writer"},
		{"ooo.stw", "application/vnd.sun.xml.writer.template"},
		{"ooo.sxw", "application/vnd.sun.xml.writer"},
		{"ooo-test.fodg", "application/vnd.oasis.opendocument.graphics-flat-xml"},
		{"ooo-test.fodp", "application/vnd.oasis.opendocument.presentation-flat-xml"},
		{"ooo-test.fods", "application/vnd.oasis.opendocument.spreadsheet-flat-xml"},
		{"ooo-test.fodt", "application/vnd.oasis.opendocument.text-flat-xml"},
		{"ooo-test.odg", "application/vnd.oasis.opendocument.graphics"},
		{"ooo-test.odp", "application/vnd.oasis.opendocument.presentation"},
		{"ooo-test.ods", "application/vnd.oasis.opendocument.spreadsheet"},
		{"ooo-test.odt", "application/vnd.oasis.opendocument.text"},
		{"ooo.vor", "application/vnd.stardivision.writer"},
		{"ooo-xp.doc", "application/msword"},
		{"Oriental_tattoo_by_daftpunk22.eps", "image/x-eps"},
		{"panasonic_lumix_dmc_fz38_05.rw2", "image/x-panasonic-rw2"},
		{"pdf-not-matlab", "application/octet-stream"},
		{"petite-ouverture-a-danser.ly", "text/x-lilypond"},
		{"pico-rom.bin", "application/x-saturn-rom"},
		{"playlist.asx", "audio/x-ms-asx"},
		{"playlist.mrl", "text/x-mrml"},
		{"playlist.wpl", "application/vnd.ms-wpl"},
		{"plugins.qmltypes", "text/x-qml"},
		{"pocket-word.psw", "application/x-pocket-word"},
		{"pom.xml", "text/x-maven+xml"},
		{"Presentation.kpt", "application/x-kpresenter"},
		{"project.glade", "application/x-glade"},
		{"PWAD.WAD", "application/x-doom-wad"},
		{"pyside.py", "text/x-python"},
		{"raw-mjpeg.mjpeg", "video/x-mjpeg"},
		{"README.pdf", "application/pdf"},
		{"rectangle.qml", "text/x-qml"},
		{"registry-nt.reg", "text/x-ms-regedit"},
		{"registry.reg", "text/x-ms-regedit"},
		{"reStructuredText.rst", "text/x-rst"},
		{"rgb-reference.ktx", "image/ktx"},
		{"ringtone.ime", "text/x-iMelody"},
		{"ringtone.m4r", "audio/x-m4r"},
		{"ringtone.mmf", "application/x-smaf"},
		{"ripoux.sap", "application/x-thomson-sap-image"},
		{"sample1.nzb", "application/x-nzb"},
		{"sample2.amr", "audio/AMR"},
		{"sample.docx", "application/vnd.openxmlformats-officedocument.wordprocessingml.document"},
		{"sample.png.lzma", "application/x-lzma"},
		{"sample.ppsx", "application/vnd.openxmlformats-officedocument.presentationml.slideshow"},
		{"sample.pptx", "application/vnd.openxmlformats-officedocument.presentationml.presentation"},
		{"sample.vsdx", "application/vnd.ms-visio.drawing.main+xml"},
		{"sample.xlsx", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"},
		{"saturn-test.bin", "application/x-saturn-rom"},
		{"SConscript", "text/x-scons"},
		{"SConscript.buildinfo", "text/x-scons"},
		{"SConstruct", "text/x-scons"},
		{"sega-cd-test.iso", "application/x-cd-image"},
		{"serafettin.rar", "application/vnd.rar"},
		{"settings.xml", "text/x-maven+xml"},
		{"settopbox.ts", "text/vnd.qt.linguist"},
		{"sg1000-test.sg", "application/x-sg1000-rom"},
		{"shebang.qml", "text/x-qml"},
		{"simon.669", "audio/x-mod"},
		{"simple-obj-c.m", "application/vnd.wolfram.mathematica.package"},
		{"small_wav.mxf", "application/mxf"},
		{"sms-test.sms", "application/x-sms-rom"},
		{"spinboxes-0.1.1-Linux.tar.xz", "application/x-xz-compressed-tar"},
		{"sqlite2.kexi", "application/x-kexiproject-sqlite2"},
		{"sqlite3.kexi", "application/x-kexiproject-sqlite2"},
		{"ssh-public-key.txt", "text/plain"},
		{"Stallman_Richard_-_The_GNU_Manifesto.fb2", "application/x-fictionbook+xml"},
		{"Stallman_Richard_-_The_GNU_Manifesto.fb2.zip", "application/x-zip-compressed-fb2"},
		{"stream.nsc", "application/x-netshow-channel"},
		{"stream.sdp", "application/sdp"},
		{"subshapes.swf", "application/vnd.adobe.flash.movie"},
		{"subtitle-microdvd.sub", "text/x-microdvd"},
		{"subtitle-mpsub.sub", "text/x-microdvd"},
		{"subtitle.smi", "application/smil+xml"},
		{"subtitle.srt", "application/x-subrip"},
		{"subtitle.ssa", "text/x-ssa"},
		{"subtitle-subviewer.sub", "text/x-microdvd"},
		{"survey.js", "application/javascript"},
		{"systemd.automount", "text/x-systemd-unit"},
		{"systemd.device", "text/x-systemd-unit"},
		{"systemd.mount", "text/x-systemd-unit"},
		{"systemd.path", "text/x-systemd-unit"},
		{"systemd.scope", "text/x-systemd-unit"},
		{"systemd.service", "text/x-dbus-service"},
		{"systemd.slice", "text/x-systemd-unit"},
		{"systemd.socket", "text/x-systemd-unit"},
		{"systemd.swap", "text/x-systemd-unit"},
		{"systemd.target", "text/x-systemd-unit"},
		{"systemd.timer", "text/x-systemd-unit"},
		{"tb-from-sentbox.eml", "message/rfc822"},
		{"tb-saved.eml", "message/rfc822"},
		{"test10.gpx", "application/gpx+xml"},
		{"test1.pcf", "application/x-cisco-vpn-settings"},
		{"test2.pcf", "application/x-cisco-vpn-settings"},
		{"test2.ppm", "image/x-portable-pixmap"},
		{"test2.tga", "image/x-tga"},
		{"test3.py", "text/x-python"},
		{"test.aa", "audio/x-pn-audibleaudio"},
		{"test.aax", "audio/x-pn-audibleaudio"},
		{"test.aiff", "audio/x-aiff"},
		{"test.alz", "application/x-alz"},
		{"test.avf", "video/x-msvideo"},
		{"test.avi", "video/x-msvideo"},
		{"test.bflng", "application/octet-stream"},
		{"test.bmp", "image/bmp"},
		{"test.bsdiff", "application/x-bsdiff"},
		{"testcase.is-really-a-pdf", "application/octet-stream"},
		{"testcases.ksp", "application/x-kspread"},
		{"test.cbl", "text/x-cobol"},
		{"test.ccmx", "application/x-ccmx"},
		{"test-cdda.toc", "application/x-cdrdao-toc"},
		{"test-cdrom.toc", "application/x-cdrdao-toc"},
		{"test.cel", "application/octet-stream"},
		{"test.cl", "text/x-opencl-src"},
		{"test.class", "application/x-java"},
		{"test.cmake", "text/x-cmake"},
		{"test.coffee", "application/vnd.coffeescript"},
		{"testcompress.z", "application/x-compress"},
		{"test.cs", "text/x-csharp"},
		{"test.csvs", "text/csv-schema"},
		{"test.d", "text/x-dsrc"},
		{"test.dcm", "application/dicom"},
		{"test.djvu", "image/vnd.djvu"},
		{"test.dot", "application/msword-template"},
		{"test.dts", "audio/vnd.dts"},
		{"test.dtshd", "audio/vnd.dts.hd"},
		{"test-en.mo", "application/x-gettext-translation"},
		{"test-en.po", "text/x-gettext-translation"},
		{"test.eps", "image/x-eps"},
		{"test.ext,v", "text/plain"},
		{"test.feature", "text/x-gherkin"},
		{"test.fit", "application/octet-stream"},
		{"test.fl", "application/x-fluid"},
		{"test.flac", "audio/flac"},
		{"test.fli", "video/x-flic"},
		{"test.gbr", "image/x-gimp-gbr"},
		{"test.gcode", "text/x.gcode"},
		{"test.geo.json", "application/geo+json"},
		{"test.geojson", "application/geo+json"},
		{"test-gettext.c", "text/x-csrc"},
		{"test.gif", "image/gif"},
		{"test.gih", "image/x-gimp-gih"},
		{"test.gnd", "application/gnunet-directory"},
		{"test.go", "text/x-go"},
		{"test.gpx", "application/gpx+xml"},
		{"test.gs", "text/x-genie"},
		{"test.h5", "application/x-hdf"},
		{"test.hdf4", "application/x-hdf"},
		{"test.hlp", "application/winhlp"},
		{"test.ico", "image/vnd.microsoft.icon"},
		{"test.ilbm", "image/x-ilbm"},
		{"test.im1", "application/octet-stream"},
		{"test.iptables", "text/x-iptables"},
		{"test.ipynb", "application/x-ipynb+json"},
		{"test.it87", "application/x-it87"},
		{"test.jar", "application/x-java-archive"},
		{"test.jceks", "application/x-java-jce-keystore"},
		{"test.jks", "application/x-java-keystore"},
		{"test.jnlp", "application/x-java-jnlp-file"},
		{"test.jpg", "image/jpeg"},
		{"test.kdc", "image/x-kodak-kdc"},
		{"test.key", "application/x-iwork-keynote-sffkey"},
		{"test-kounavail2.kwd", "application/x-kword"},
		{"test.lwp", "application/vnd.lotus-wordpro"},
		{"test.lz", "application/x-lzip"},
		{"test.lz4", "application/x-lz4"},
		{"test.lzo", "application/x-lzop"},
		{"test.manifest", "text/cache-manifest"},
		{"test.metalink", "application/metalink+xml"},
		{"test.mml", "application/mathml+xml"},
		{"test.mng", "video/x-mng"},
		{"test.mo", "application/x-gettext-translation"},
		{"test.mobi", "application/x-mobipocket-ebook"},
		{"test.mof", "text/x-mof"},
		{"test.msi", "application/x-msi"},
		{"test-noid3.mp3", "audio/mpeg"},
		{"test.ogg", "audio/ogg"},
		{"test.ooc", "text/x-ooc"},
		{"test.opus", "audio/ogg"},
		{"test.owx", "application/owl+xml"},
		{"test.p12", "application/pkcs12"},
		{"test-p6.ppm", "image/x-portable-pixmap"},
		{"test.p7b", "application/pkcs7-mime"},
		{"test.pat", "image/x-gimp-pat"},
		{"test.pbm", "image/x-portable-bitmap"},
		{"test.pcx", "image/vnd.zbrush.pcx"},
		{"test.pdf.lz", "application/x-lzpdf"},
		{"test.pdf.xz", "application/x-xzpdf"},
		{"test.pgm", "image/x-portable-graymap"},
		{"test.pgn", "application/vnd.chess-pgn"},
		{"test.php", "application/x-php"},
		{"test.pix", "application/octet-stream"},
		{"test.pkipath", "application/pkix-pkipath"},
		{"test.pl", "application/x-perl"},
		{"test.pm", "application/x-pagemaker"},
		{"test.pmd", "application/x-pagemaker"},
		{"test.png", "image/png"},
		{"test.por", "application/x-spss-por"},
		{"test.pot", "application/vnd.ms-powerpoint"},
		{"test.ppm", "image/x-portable-pixmap"},
		{"test.ps", "application/postscript"},
		{"test.psd", "image/vnd.adobe.photoshop"},
		{"test-public-key.asc", "text/plain"},
		{"test.py", "text/x-python"},
		{"test.py3", "text/x-python3"},
		{"test.pyx", "text/x-python"},
		{"test.qp", "application/x-qpress"},
		{"test.qti", "application/x-qtiplot"},
		{"test.raml", "application/raml+yaml"},
		{"test.random", "application/octet-stream"},
		{"test-reordered.ipynb", "application/x-ipynb+json"},
		{"test.rs", "text/rust"},
		{"test.sass", "text/x-sass"},
		{"test.sav", "application/x-spss-sav"},
		{"test.scala", "text/x-scala"},
		{"test.scm", "text/x-scheme"},
		{"test.scss", "text/x-scss"},
		{"test-secret-key.asc", "text/plain"},
		{"test-secret-key.skr", "application/pgp-keys"},
		{"test.sgi", "image/x-sgi"},
		{"test.sqlite2", "application/x-sqlite2"},
		{"test.sqlite3", "application/vnd.sqlite3"},
		{"test.ss", "text/x-scheme"},
		{"test.sv", "text/x-svsrc"},
		{"test.svh", "text/x-svhdr"},
		{"test.t", "application/x-perl"},
		{"test.tar.lz", "application/x-lzip-compressed-tar"},
		{"test.tar.lz4", "application/x-lz4-compressed-tar"},
		{"test-template.dot", "application/msword-template"},
		{"test.tex", "text/x-tex"},
		{"test.tga", "image/x-tga"},
		{"test.tif", "image/tiff"},
		{"test.ts", "text/vnd.qt.linguist"},
		{"test.ttl", "text/turtle"},
		{"test.ttx", "application/x-font-ttx"},
		{"test.twig", "text/x-twig"},
		{"test.url", "application/x-mswinurl"},
		{"test.uue", "text/x-uuencode"},
		{"test.v", "text/x-verilog"},
		{"test.vala", "text/x-vala"},
		{"test.vcf", "text/vcard"},
		{"test-vpn.pcf", "application/x-cisco-vpn-settings"},
		{"test.vsd", "application/vnd.visio"},
		{"test.wav", "audio/x-wav"},
		{"test.webp", "image/webp"},
		{"test.wim", "application/x-ms-wim"},
		{"test.wps", "application/vnd.ms-works"},
		{"test.xar", "application/x-xar"},
		{"test.xbm", "image/x-xbitmap"},
		{"test.xcf", "image/x-xcf"},
		{"test.xht", "application/xhtml+xml"},
		{"test.xhtml", "application/xhtml+xml"},
		{"test.xlr", "application/vnd.ms-works"},
		{"test.xml.in", "application/octet-stream"},
		{"test.xpm", "image/x-xpixmap"},
		{"test.xsl", "application/xslt+xml"},
		{"test.xwd", "image/x-xwindowdump"},
		{"test.yaml", "application/x-yaml"},
		{"test.zip", "application/zip"},
		{"test.zz", "application/zlib"},
		{"text-iso8859-15.txt", "text/plain"},
		{"text.pdf", "application/pdf"},
		{"text.ps", "application/postscript"},
		{"text.ps.gz", "application/x-gzpostscript"},
		{"text.PS.gz", "application/x-gzpostscript"},
		{"text.qmlproject", "text/x-qml"},
		{"text-utf8.txt", "text/plain"},
		{"text.wwf", "application/x-wwf"},
		{"tree-list", "application/octet-stream"},
		{"TS010082249.pub", "application/vnd.ms-publisher"},
		{"upc-video-subtitles-en.vtt", "text/vtt"},
		{"Utils.jsm", "application/javascript"},
		{"virtual-boy-wario-land.vb", "application/x-virtual-boy-rom"},
		{"weather_sun.xcf", "image/x-xcf"},
		{"webfinger.jrd", "application/jrd+json"},
		{"white_640x480.kra", "application/x-krita"},
		{"wii.wad", "application/x-doom-wad"},
		{"wonderswan-color-chocobo.wsc", "application/x-wonderswan-color-rom"},
		{"wonderswan-rockman-forte.ws", "application/x-wonderswan-rom"},
		{"xml-in-mp3.mp3", "audio/mpeg"},
	}
	for _, test := range tests {
		t.Run(test.filename, func(t *testing.T) {
			if got := MatchGlob(test.filename).MediaType(); got != test.want {
				t.Errorf("MatchGlob() = %v, want %v", got, test.want)
			}
		})
	}
	prefixesCS["test."] = []simpleGlob{{90, 0}}
	want := "all/all"
	filename := "test.file"
	t.Run(filename, func(t *testing.T) {
		if got := MatchGlob(filename).MediaType(); got != want {
			t.Errorf("MatchGlob() = %v, want %v", got, want)
		}
	})
	prefixesCS["test."] = nil
	globs = append(globs, prefixPattern{pattern{
		matchers: []matcher{list("tvx"), list("wey"), list("spk"), list("wmt"), value(".")},
		length:   5,
	}, true, 1, 100})
	want = "all/allfiles"
	t.Run(filename, func(t *testing.T) {
		if got := MatchGlob(filename).MediaType(); got != want {
			t.Errorf("MatchGlob() = %v, want %v", got, want)
		}
	})
	globs = globs[:len(globs)-1]
}

func benchmarkMatchGlob(filename string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		MatchGlob(filename)
	}
}

func BenchmarkMatchGlob(b *testing.B) {
	for _, f := range combinedTests {
		b.Run(f.filename, func(b *testing.B) {
			benchmarkMatchGlob(f.filename, b)
		})
	}
}
