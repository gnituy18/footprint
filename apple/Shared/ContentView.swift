//
//  ContentView.swift
//  Shared
//
//  Created by Charlie Hsu on 2021/2/5.
//

import SwiftUI

struct ContentView: View {
    var body: some View {
		Text(config.backendURL)
            .padding()
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
    }
}
