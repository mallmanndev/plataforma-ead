ffmpeg \
    -i input.mp4 \
    -hls_time 10 \
    -hls_list_size 0 \
    -hls_segment_filename "resolucao-1080/segment%d.ts" \
    -f hls resolucao-1080/index.m3u8
