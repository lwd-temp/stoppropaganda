package targets

// Source: https://twitter.com/FedorovMykhailo/status/1497642156076511233

var TargetWebsites = map[string]struct{}{
	/* Other countries */

	"https://bukimevieningi.lt": {},
	"https://musutv.lt":         {},
	"https://api.musutv.lt":     {},
	"https://lt.baltnews.com":   {},
	"https://lt.rubaltic.ru":    {},
	"http://sputniknews.lt":     {},
	"https://lv.sputniknews.ru": {},
	"https://viada.lt":          {},
	"https://api.viada.lt":      {},
	"https://www.sber.kz":       {},
	"https://www.sberbank.kz":   {},

	/* Russia */

	// Propaganda
	"https://lenta.ru":                    {},
	"https://ria.ru":                      {},
	"https://ria.ru/lenta":                {},
	"https://www.rbc.ru":                  {},
	"https://smotrim.ru":                  {},
	"https://api.smotrim.ru":              {},
	"https://tass.ru":                     {},
	"https://tass.ru/userApi/getNewsFeed": {},
	"https://tvzvezda.ru":                 {},
	"https://www.vesti.ru":                {},
	"https://zakupki.gov.ru":              {},
	"https://er.ru":                       {},
	"https://www.rzd.ru":                  {},
	"https://rzdlog.ru":                   {},
	"https://vgtrk.ru":                    {},
	"https://www.interfax.ru":             {},
	"https://iz.ru":                       {},
	"https://vz.ru":                       {},
	"https://sputniknews.ru":              {},
	"https://www.gazeta.ru":               {},
	"https://www.kp.ru":                   {},
	"https://riafan.ru":                   {},
	"https://api.riafan.ru":               {},
	"https://pikabu.ru":                   {},
	"https://api.pikabu.ru":               {},
	"https://www.kommersant.ru":           {},
	"https://omk.ru":                      {},
	"https://regnum.ru":                   {},
	"https://www.rambler.ru":              {},
	"https://mail.ru":                     {},
	"https://www.astrobl.ru":              {},
	"https://www.volgograd.ru":            {},
	"https://www.admoblkaluga.ru":         {},
	"https://apparat.lenobl.ru":           {},
	"https://mosreg.ru":                   {},
	"https://orel-region.ru":              {},
	"https://midural.ru":                  {},
	"https://ulgov.ru":                    {},
	"https://cheladmin.ru":                {},
	"https://www.stavregion.ru":           {},
	"https://www.mk.ru":                   {},
	"https://rg.ru":                       {},

	// Business corporations
	"https://lukoil.ru":                  {},
	"https://magnit.ru":                  {},
	"https://www.evraz.com/ru":           {},
	"https://nlmk.com":                   {},
	"https://www.sibur.ru":               {},
	"https://www.severstal.com":          {},
	"https://www.eurosib.ru":             {},
	"https://www.wildberries.ru":         {},
	"https://www.ozon.ru":                {},
	"https://www.dns-shop.ru":            {},
	"https://aliexpress.ru":              {},
	"https://advego.com":                 {},
	"https://freelance.ru":               {},
	"https://www.turbotext.ru":           {},
	"https://sberfn.ru":                  {},
	"https://sber-am.ru":                 {},
	"https://www.vtbcapital-pr.ru":       {},
	"https://region-am.ru":               {},
	"https://www.ingosinvest.ru":         {},
	"https://goszakaz.ru":                {},
	"https://star-pro.ru":                {},
	"https://region.ru":                  {},
	"http://gruzovozkin.pro":             {},
	"https://ok.ru":                      {},
	"http://5.61.23.11":                  {},
	"http://217.20.155.13":               {},
	"http://217.20.147.1":                {},
	"http://www.yemelya.ru":              {},
	"https://www.nornickel.com":          {},
	"https://www.evraz.com":              {},
	"https://rostec.ru":                  {},
	"https://scloud.rostec.ru/login":     {},
	"https://vcs.rostec.ru":              {},
	"https://lk.rostec.ru/Account/LogIn": {},
	"http://ias.rostec.ru":               {},
	"https://smi.rostec.ru/user":         {},
	"https://vks3.rostec.ru":             {},
	"https://kontur.ru":                  {},
	"https://help.kontur.ru/vnc":         {},
	"https://geltek-medica.ru":           {},

	// Banks
	"https://www.sberbank.ru":                                        {},
	"https://online.sberbank.ru":                                     {},
	"https://api.developer.sber.ru/product/SberbankID":               {},
	"https://api.sberbank.ru/prod/tokens/v2":                         {},
	"https://api.sberbank.ru/prod/tokens/v2/oauth":                   {},
	"https://api.sberbank.ru/prod/tokens/v2/oidc":                    {},
	"https://www.vtb.ru":                                             {},
	"https://api.vtb.ru":                                             {},
	"https://www.gazprombank.ru":                                     {},
	"https://api.gazprombank.ru":                                     {},
	"https://www.moex.com":                                           {},
	"https://iss.moex.com/iss/reference":                             {},
	"https://messaging.moex.com/init":                                {},
	"https://passport.moex.com/login":                                {},
	"http://www.fsb.ru":                                              {},
	"https://scr.online.sberbank.ru/api/fl/idgib-w-3ds":              {},
	"https://3dsec.sberbank.ru/mportal3/auth/login":                  {},
	"https://acs1.sbrf.ru":                                           {},
	"https://acs2.sbrf.ru":                                           {},
	"https://acs3.sbrf.ru":                                           {},
	"https://acs4.sbrf.ru":                                           {},
	"https://acs5.sbrf.ru":                                           {},
	"https://acs6.sbrf.ru":                                           {},
	"https://acs7.sbrf.ru":                                           {},
	"https://acs8.sbrf.ru":                                           {},
	"https://my.bank-hlynov.ru":                                      {},
	"https://chbrr.crimea.com":                                       {},
	"https://enter.unicredit.ru/v2/cgi/bsi.dll?T=RT_2Auth.BF":        {},
	"https://online.vtb.ru":                                          {},
	"https://online.sberbank.ru/CSAFront/index.do":                   {},
	"https://online.gpb.ru/login":                                    {},
	"https://alfabank.ru":                                            {},
	"https://alfabank.ru/api/v1/geco-ip":                             {}, // throws 500. flood their loggers
	"https://rshb.ru":                                                {},
	"https://online.rshb.ru/b1/ib6/wf2/retail/ib/loginretaildefault": {},
	"https://online.sovcombank.ru":                                   {},
	"https://online.mkb.ru":                                          {},
	"https://www.tinkoff.ru":                                         {},
	"https://id.tinkoff.ru/auth/step":                                {},
	"https://178.248.236.218":                                        {},
	"https://91.194.226.50":                                          {},
	"https://auth.tcsbank.ru":                                        {},
	"https://91.194.226.32":                                          {},

	//The state
	"https://www.gosuslugi.ru":                 {},
	"https://www.gosuslugi.ru/api/mainpage/v4": {},
	"https://www.mos.ru/uslugi":                {},
	"https://api.mos.ru":                       {},
	"http://kremlin.ru":                        {},
	"http://government.ru":                     {},
	"https://mil.ru":                           {},
	"https://www.nalog.gov.ru":                 {},
	"https://customs.gov.ru":                   {},
	"https://pfr.gov.ru":                       {},
	"https://rkn.gov.ru":                       {},
	"https://gosuslugi41.ru":                   {},
	"https://uslugi27.ru":                      {},
	"https://gosuslugi29.ru":                   {},
	"https://gosuslugi.astrobl.ru":             {},
	"http://pochta.ru":                         {},
	"http://crimea-post.ru":                    {},
	"https://ca.vks.rosguard.gov.ru":           {},

	// Embassies
	// Do not duplicate the IPs: https://github.com/erkexzcx/stoppropaganda/pull/110#issuecomment-1059960305
	"https://montreal.mid.ru": {},
	"https://belarus.mid.ru":  {},

	// Others
	"https://mail.rkn.gov.ru":        {},
	"https://cloud.rkn.gov.ru":       {},
	"https://mvd.gov.ru":             {},
	"https://pwd.wto.economy.gov.ru": {},
	"https://stroi.gov.ru":           {},
	"https://proverki.gov.ru":        {},
	"https://www.glonass-iac.ru":     {},
	"https://savelife.pw":            {},
	// former russian project with russian founders, that are being used as navigational service for russian army
	// has marked military objects
	"https://wikimapia.org": {},
	"https://dme.ru":        {},

	// Payments
	"http://185.170.2.7":                  {},
	"http://185.170.3.7":                  {},
	"http://185.170.2.9":                  {},
	"https://ds1.mirconnect.ru":           {},
	"https://ds2.mirconnect.ru":           {},
	"https://uat-ds1.mirconnect.ru":       {},
	"https://uat-ds2.mirconnect.ru":       {},
	"https://sberpay.mirconnect.ru":       {},
	"https://wats.mirconnect.ru":          {},
	"https://ds.vendorcert.mirconnect.ru": {},
	"https://privetmir.ru":                {},
	"https://mironline.ru":                {},
	"https://sbp.nspk.ru":                 {},
	"https://nspk.ru":                     {},
	"https://nspk.com":                    {},
	"https://sip.nspk.ru":                 {},
	"https://koronapay.com":               {},
	"https://api.koronapay.com":           {},
	"https://etpgpb.ru":                   {},
	"https://api.etpgpb.ru":               {},
	"https://etp.gpb.ru":                  {},
	"https://passport.etpgpb.ru":          {},
	"https://gos.etpgpb.ru":               {},
	"https://inter-rao.etpgpb.ru":         {},
	"https://rosnedra.etpgpb.ru":          {},
	"https://gpn.etpgpb.ru":               {},
	"https://geh.etpgpb.ru":               {},
	"https://trade.etpgpb.ru":             {},
	"https://gb.etpgpb.ru":                {},
	"https://bki-okb.ru":                  {},
	"http://185.215.4.14":                 {},
	"http://185.215.4.10":                 {},
	"https://185.9.230.77":                {},
	"https://195.10.198.201":              {},
	"https://195.10.198.136":              {},
	"https://195.10.198.243":              {},
	"https://195.10.198.204":              {},
	"https://195.10.198.208":              {},

	// Exchanges connected to russian banks
	"https://betatransfer.org":        {},
	"https://bitokk.biz":              {},
	"https://www.netex24.net":         {},
	"https://flashobmen.com":          {},
	"https://ychanger.net":            {},
	"https://multichange.net":         {},
	"https://royal.cash":              {},
	"https://prostocash.com":          {},
	"https://baksman.org":             {},
	"https://yoomoney.ru":             {},
	"https://yookassa.ru":             {},
	"https://telemetry.koronapay.com": {},
	"https://olympus.koronapay.com":   {},

	// Electronic signature services, certificate authorities, www domain names
	"http://www.nucrf.ru":                {},
	"http://www.belinfonalog.ru":         {},
	"https://www.roseltorg.ru":           {},
	"https://178fz.roseltorg.ru":         {},
	"https://agro.roseltorg.ru":          {},
	"https://atom2.roseltorg.ru":         {},
	"https://bg.roseltorg.ru/auth/email": {},
	"https://com.roseltorg.ru":           {},
	"https://docs.roseltorg.ru":          {},
	"https://etp.roseltorg.ru":           {},
	"https://fkr.roseltorg.ru":           {},
	"https://fkr2.roseltorg.ru":          {},
	"https://kim-atom.roseltorg.ru":      {},
	"https://kim.roseltorg.ru":           {},
	"https://kim-irao.roseltorg.ru":      {},
	"https://lk.roseltorg.ru/auth":       {},
	"https://orders.roseltorg.ru":        {},
	"https://rosgeo.roseltorg.ru":        {},
	"https://rosseti.roseltorg.ru":       {},
	"https://rt.roseltorg.ru":            {},
	"https://rushydro.roseltorg.ru":      {},
	"https://msp.roseltorg.ru":           {},
	"https://vtb.roseltorg.ru":           {},
	"https://rosinvoice.ru/auth":         {},
	"https://astral.ru":                  {},
	"http://www.nwudc.ru":                {},
	"https://kk.bank":                    {},
	"https://structure.mil.ru":           {},
	"http://www.ucpir.ru":                {},
	"http://www.e-portal.ru":             {},
	"https://api.e-portal.ru":            {},
	"https://parus-s.ru":                 {},
	"http://www.icentr.ru":               {},
	"https://www.kartoteka.ru":           {},
	"https://api.kartoteka.ru":           {},
	"http://rsbis.ru":                    {},
	"http://www.stv-it.ru":               {},
	"https://crypset.ru":                 {},
	"http://www.24ecp.ru":                {},
	"http://kraskript.com":               {},
	"http://ca.ntssoft.ru":               {},
	"http://www.y-center.ru":             {},
	"http://www.rcarus.ru":               {},
	"http://squaretrade.ru":              {},
	"http://ca.gisca.ru":                 {},
	"http://www.otchet-online.ru":        {},
	"http://udcs.ru":                     {},
	"http://www.cit-ufa.ru":              {},
	"http://api.cit-ufa.ru":              {},
	"http://elkursk.ru":                  {},
	"http://www.icvibor.ru":              {},
	"http://ucestp.ru":                   {},
	"https://api.ucestp.ru":              {},
	"http://mcspro.ru":                   {},
	"http://www.infotrust.ru":            {},
	"http://epnow.ru":                    {},
	"http://cfmc.ru":                     {},
	"https://esia.gosuslugi.ru":          {},
	"https://iecp.ru/ep/ep-verification": {},
	"https://e-trust.gosuslugi.ru":       {},
	"https://gu.spb.ru":                  {},
	"https://iecp.ru/ep/uc-list":         {},
	"http://glovo.ru":                    {},

	// Oil and gas trading companies
	"https://www.tektorg.ru": {},
	"https://rosneft.com":    {},
	"https://lukoil.com":     {},
	"https://aaa-srt-cs.lukoil.com/logon/LogonPoint/tmindex.html": {},
	"https://aaa-srt-apps.lukoil.com":                             {},
	"https://gazprom.ru":                                          {},
	"https://b2b.sibur.ru":                                        {},
	"https://onlinecontract.ru":                                   {},
	"https://www.eurochemgroup.com":                               {},
	"https://uralchem.com":                                        {},
	"https://tatneft.ru":                                          {},
	"https://acron.ru":                                            {},

	// Food delivery services
	"https://www.delivery-club.ru": {},
	"https://m-food.ru":            {},
	"https://sbermarket.ru":        {},
	"https://chibbis.ru":           {},
	"https://samokat.ru":           {},
	"https://localkitchen.ru":      {},

	// Cinemas
	"https://cinemastar.ru":     {},
	"https://kinoteatr.ru":      {},
	"https://karofilm.ru":       {},
	"https://kinomax.ru":        {},
	"https://romanov-cinema.ru": {},
	"https://pioner-cinema.ru":  {},
	"https://www.mirage.ru":     {},

	// Others
	"http://217.12.104.100": {},
	"http://172.217.0.179":  {},
	"http://193.104.87.172": {},
	"http://213.33.173.184": {},

	// "Internet and thematic forums"
	"https://www.eapteka.ru":    {},
	"https://www.asna.ru":       {},
	"https://366.ru":            {},
	"https://aptekamos.ru":      {},
	"https://vseapteki.ru":      {},
	"https://www.rigla.ru":      {},
	"https://planetazdorovo.ru": {},
	"https://samson-pharma.ru":  {},
	"https://zdorov.ru":         {},
	"https://apteka.ru":         {},

	// Various websites by ip
	"https://46.28.16.112":   {},
	"https://91.213.144.193": {},
	"https://91.213.144.237": {},
	"https://178.248.238.24": {},
	"https://212.24.38.190":  {},
	"https://178.57.94.11":   {},
	"https://185.10.60.69":   {},
	"https://185.10.60.88":   {},
	"https://78.47.115.99":   {},
	"https://81.19.72.39":    {},
	"https://81.19.72.3":     {},
	"https://194.190.37.226": {},
	"https://194.190.37.228": {},
	"https://194.190.37.234": {},
	"https://213.59.254.8":   {},
	"https://185.194.34.142": {},
	"https://213.59.197.65":  {},
	"https://185.71.67.237":  {},
	"https://82.151.111.187": {},
	"https://91.239.5.18":    {},
	"https://91.239.5.75":    {},
	"https://185.58.223.206": {},
	"https://178.248.238.15": {},
	"https://92.53.98.191":   {},
	"https://37.230.115.34":  {},
	"https://213.159.213.33": {},

	/* BELARUS */

	// by gov
	"https://mininform.gov.by":       {},
	"https://rec.gov.by/ru":          {},
	"https://www.mil.by":             {},
	"https://www.government.by":      {},
	"https://president.gov.by/ru":    {},
	"https://www.mvd.gov.by/ru":      {},
	"http://www.kgb.by/ru":           {},
	"https://www.prokuratura.gov.by": {},
	"http://mfa.gov.by":              {},
	"https://russia.mfa.gov.by":      {},

	// by banks
	"https://www.nbrb.by":                 {},
	"https://belarusbank.by":              {},
	"https://brrb.by":                     {},
	"https://www.belapb.by":               {},
	"https://bankdabrabyt.by":             {},
	"https://belinvestbank.by/individual": {},
	"https://api.belinvestbank.by":        {},
	"https://belpost.by":                  {},

	// by business
	"https://bgp.by/ru":           {},
	"https://www.belneftekhim.by": {},
	"http://www.bellegprom.by":    {},
	"https://www.energo.by":       {},
	"http://belres.by/ru":         {},
	"https://www.rw.by":           {},

	// by media
	"https://belta.by":            {},
	"https://sputnik.by":          {},
	"https://www.tvr.by":          {},
	"https://www.sb.by":           {},
	"https://www.belarus.by":      {},
	"https://belarus24.by":        {},
	"https://ont.by":              {},
	"https://www.024.by":          {},
	"https://www.belnovosti.by":   {},
	"https://mogilevnews.by":      {},
	"https://yandex.by":           {},
	"https://www.slonves.by":      {},
	"http://www.ctv.by":           {},
	"https://radiobelarus.by":     {},
	"https://radiusfm.by":         {},
	"https://alfaradio.by":        {},
	"https://radiomir.by":         {},
	"https://api.radiomir.by":     {},
	"https://radiostalica.by":     {},
	"https://radiobrestfm.by":     {},
	"https://api.radiobrestfm.by": {},
	"https://www.tvrmogilev.by":   {},
	"https://zarya.by":            {},
	"https://grodnonews.by":       {},

	/* Syria */
	"https://syrianfinance.gov.sy": {},
	"http://185.216.132.201":       {},

	/* DDOS mitigation */
	"https://ddos-guard.net/ru": {},
	"https://stormwall.pro":     {},
	"https://qrator.net/ru":     {},
	"https://solidwall.ru":      {},

	// Still operating in Russia
	// https://yale.box.com/s/11lqy1d3yn1kf9xa3r96k9sb6w5m4qea
	"https://www.auchan.ru": {},
	"https://www.blablacar.ru": {},
	"https://auth.blablacar.ru": {},
	"https://edge.blablacar.ru/location/suggestions": {},
	"https://burgerkingrus.ru": {},
	"https://burgerkingrus.ru/api-web-front/api/v3/restaurant/list": {},
	"https://mic.burgerking.ru/mifs/user/login.jsp": {},
	"https://sd.burgerking.ru/HEAT/SaaSExternalSessionRenew.aspx": {},
	"https://www.decathlon.ru": {},
	"https://ecco.ru": {},
	"https://www.leroymerlin.ru": {},
	"https://mega.ru/online/": {},
	"https://mega.ru/online/rest/catalog/list/favourite/2557159/": {},  // Add to favs via get
	"https://papajohns.ru/moscow": {},
	"https://api.papajohns.ru/slider/list": {},
	"https://www.renault.ru": {},
	"https://subway.ru": {},
	"https://www.yves-rocher.ru": {},

    // https://t.me/itarmyofukraine2022/230
	"http://www.ved.gov.ru":      {},
	"https://www.mid.ru":         {},
	"https://www.economy.gov.ru": {},
	"https://minobrnauki.gov.ru": {},

    // https://t.me/itarmyofukraine2022/216
	"https://belqi.net": {},

    // https://t.me/itarmyofukraine2022/215
	"https://www.chechnya.online": {},
	"https://vaynahavia.com":      {},
	"https://grozmer.ru":          {},
	"https://www.minfinchr.ru":    {},

	"https://rostelecom.ru":     {},
	"https://www.company.rt.ru": {},

    // https://t.me/itarmyofukraine2022/211
	"https://topcor.ru": {},

	// https://t.me/itarmyofukraine2022/243
	"http://fss.gov.ru":               {},
	"http://fss.ru":                   {},
	"http://forum.fss.ru":             {},
	"http://wiki.fss.ru":              {},
	"https://autodiscover.fss.ru/owa": {},
	"https://portal.fss.ru":           {},
	"https://data.fss.ru/open":        {},
	"http://docs.fss.ru":              {},
	"https://beta.fss.ru":             {},
	"https://pfrf.ru":                 {},
	"https://hosting.pfrf.ru":         {},
	"https://school.pfrf.ru":          {},
	"https://es.pfrf.ru":              {},

	// https://t.me/itarmyofukraine2022/246
	"https://rosleshoz.gov.ru": {},

	// https://t.me/itarmyofukraine2022/247
	"https://qiwi.com":          {},
	"https://checkout.qiwi.com": {},
	"https://my.qiwi.com":       {},
	"https://oplata.qiwi.com":   {},
	"https://p2p.qiwi.com":      {},

	// https://t.me/itarmyofukraine2022/248
	"https://digital.gov.ru":      {},
	"https://minenergo.gov.ru":    {},
	"https://minfin.gov.ru/ru":    {},
	"https://minjust.gov.ru":      {},
	"https://www.minsport.gov.ru": {},
	"https://minstroyrf.gov.ru":   {},
	"https://mintrud.gov.ru":      {},

	// https://t.me/itarmyofukraine2022/251
	"https://payanyway.ru": {},

	// https://t.me/itarmyofukraine2022/253
	"https://hh.ru":           {},
	"https://www.superjob.ru": {},
	"https://www.zarplata.ru": {},
	"https://rabota.vk.com":   {},
	"https://www.avito.ru":    {},
	"https://m.avito.ru":      {},
}
