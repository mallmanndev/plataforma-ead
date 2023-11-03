package integration_test

import (
	"testing"
)

func TestVideoUpload(t *testing.T) {
	t.Skip()
	/*
		db, disconnect := testutils.DatabaseConnection()
		defer disconnect()

		t.Run("When file type is invalid, Returns error", func(t *testing.T) {
			// GIVEN
			ctx, client, closer := testutils.FilesServer(db)
			defer closer()

			stream, err := client.VideoUpload(ctx)
			assert.NoError(t, err)

			reqInfo := &pb.VideoUploadRequest{
				Data: &pb.VideoUploadRequest_Info{
					Info: &pb.VideoInfo{Type: "invalid", Size: 1000},
				},
			}

			// WHEN
			err = stream.Send(reqInfo)
			assert.NoError(t, err)

			// THEN
			res, err := stream.CloseAndRecv()
			s, _ := status.FromError(err)
			assert.Nil(t, res)
			assert.Equal(t, "InvalidArgument", s.Code().String())
			assert.ErrorContains(t, err, "[Video] Invalid 'type': must be mp4.")
		})

		t.Run("When not send chunks, Returns error", func(t *testing.T) {
			// GIVEN
			ctx, client, closer := testutils.FilesServer(db)
			defer closer()

			stream, err := client.VideoUpload(ctx)
			assert.NoError(t, err)

			reqInfo := &pb.VideoUploadRequest{
				Data: &pb.VideoUploadRequest_Info{
					Info: &pb.VideoInfo{Type: "mp4", Size: 0},
				},
			}

			// WHEN
			err = stream.Send(reqInfo)
			assert.NoError(t, err)

			// THEN
			res, err := stream.CloseAndRecv()
			s, _ := status.FromError(err)
			assert.Nil(t, res)
			assert.Equal(t, "InvalidArgument", s.Code().String())
			assert.ErrorContains(t, err, "[Video Upload] Invalid 'size': Expected 0 bytes, Received 0 bytes.")
		})

		t.Run("When chunks is lower to file size, Returns error", func(t *testing.T) {
			// GIVEN
			file, err := os.Open("../fixtures/videos/720p.mp4")
			assert.NoError(t, err)
			defer file.Close()

			ctx, client, closer := testutils.FilesServer(db)
			defer closer()

			stream, err := client.VideoUpload(ctx)
			assert.NoError(t, err)

			reqInfo := &pb.VideoUploadRequest{
				Data: &pb.VideoUploadRequest_Info{
					Info: &pb.VideoInfo{Type: "mp4", Size: 2000000},
				},
			}

			// WHEN
			stream.Send(reqInfo)

			oneKb := 1024
			buffer := make([]byte, oneKb)

			for {
				// Leia um chunk do arquivo
				_, err := file.Read(buffer)
				if err != nil {
					if err.Error() == "EOF" {
						break // Fim do arquivo
					}
					t.Errorf("Não foi possível ler o arquivo: %v", err)
				}
				stream.Send(
					&pb.VideoUploadRequest{
						Data: &pb.VideoUploadRequest_Chunk{
							Chunk: buffer,
						},
					},
				)
			}

			// THEN
			res, err := stream.CloseAndRecv()
			s, _ := status.FromError(err)
			assert.Nil(t, res)
			assert.Equal(t, "InvalidArgument", s.Code().String())
			assert.ErrorContains(t, err, "[Video Upload] Invalid 'size': Expected 2000000 bytes, Received 1055744 bytes.")
		})

		t.Run("When send chunks, Returns video", func(t *testing.T) {
			// GIVEN

			// Open file
			file, err := os.Open("../fixtures/videos/720p.mp4")
			assert.NoError(t, err)
			fi, _ := file.Stat()
			fileZize := fi.Size()
			defer file.Close()

			// Start gRPC server
			ctx, client, closer := testutils.FilesServer(db)
			defer closer()

			// Create gRPC client
			stream, err := client.VideoUpload(ctx)
			assert.NoError(t, err)

			reqInfo := &pb.VideoUploadRequest{
				Data: &pb.VideoUploadRequest_Info{
					Info: &pb.VideoInfo{Type: "mp4", Size: fileZize},
				},
			}

			// WHEN
			stream.Send(reqInfo)

			oneKb := 1024
			buffer := make([]byte, oneKb)

			for {
				// Leia um chunk do arquivo
				_, err := file.Read(buffer)
				if err != nil {
					if err.Error() == "EOF" {
						break
					}
					t.Errorf("Não foi possível ler o arquivo: %v", err)
				}
				stream.Send(
					&pb.VideoUploadRequest{
						Data: &pb.VideoUploadRequest_Chunk{
							Chunk: buffer,
						},
					},
				)
			}

			// THEN
			res, err := stream.CloseAndRecv()
			fmt.Println(res)
			assert.Nil(t, err)
			if assert.NotNil(t, res) {
				assert.NotEmpty(t, res.Id)

				_, err = os.Stat(fmt.Sprintf("/videos/tmp/%s.%s", res.Id, "mp4"))
				if err != nil {
					t.Fatalf("O arquivo não foi criado: %v", err)
				}
			}
		})
	*/
}
