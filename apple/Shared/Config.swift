//
//  Config.swift
//  Footprint
//
//  Created by sfhp on 2021/3/11.
//

import Foundation

let config = getConfig()

struct Config: Decodable {
	private(set) var backendURL: String
}

private func getConfig() -> Config {
	let config = Bundle.main.infoDictionary!["Config"] as! [String: Any]
	return Config(
		backendURL: config["backendURL"] as! String
	)
}
