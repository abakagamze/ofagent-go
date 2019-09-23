/*
   Copyright 2017 the original author or authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

        http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package grpc

import (
	pb "github.com/opencord/voltha-protos/go/voltha"
	"google.golang.org/grpc"
	"log"
)
import "context"

func streamPacketOut(client pb.VolthaServiceClient) {
	opt := grpc.EmptyCallOption{}
	outClient, err := client.StreamPacketsOut(context.Background(), opt)
	if err != nil {

		log.Printf("Error creating packetout stream %v", err)
	}
	packetOutChannel := make(chan pb.PacketOut)
	for {
		ofPacketOut := <-packetOutChannel
		outClient.Send(&ofPacketOut)
	}

}
