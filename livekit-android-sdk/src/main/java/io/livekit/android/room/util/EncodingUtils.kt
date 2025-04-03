/*
 * Copyright 2023-2024 LiveKit, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package io.livekit.android.room.util

import io.livekit.android.room.track.VideoEncoding
import io.livekit.android.room.track.VideoPreset
import io.livekit.android.room.track.VideoPreset169
import io.livekit.android.room.track.VideoPreset43
import io.livekit.android.room.track.video.ScalabilityMode
import livekit.LivekitModels
import livekit.org.webrtc.RtpParameters
import kotlin.math.abs
import kotlin.math.ceil
import kotlin.math.max
import kotlin.math.min
import kotlin.math.pow
import kotlin.math.roundToInt

internal object EncodingUtils {

    val VIDEO_RIDS = arrayOf("q", "h", "f")

    // Note: maintain order from smallest to biggest.
    private val PRESETS_16_9 = listOf(
        VideoPreset169.H90,
        VideoPreset169.H180,
        VideoPreset169.H216,
        VideoPreset169.H360,
        VideoPreset169.H540,
        VideoPreset169.H720,
        VideoPreset169.H1080,
        VideoPreset169.H1440,
        VideoPreset169.H2160,
    )

    // Note: maintain order from smallest to biggest.
    private val PRESETS_4_3 = listOf(
        VideoPreset43.H120,
        VideoPreset43.H180,
        VideoPreset43.H240,
        VideoPreset43.H360,
        VideoPreset43.H480,
        VideoPreset43.H540,
        VideoPreset43.H720,
        VideoPreset43.H1080,
        VideoPreset43.H1440,
    )

    fun determineAppropriateEncoding(width: Int, height: Int): VideoEncoding {
        val presets = presetsForResolution(width, height)

        // presets assume width is longest size
        val longestSize = max(width, height)
        val preset = presets
            .firstOrNull { it.capture.width >= longestSize }
            ?: presets.last()

        return preset.encoding
    }

    fun presetsForResolution(width: Int, height: Int): List<VideoPreset> {
        val longestSize = max(width, height)
        val shortestSize = min(width, height)
        val aspectRatio = longestSize.toFloat() / shortestSize
        return if (abs(aspectRatio - 16f / 9f) < abs(aspectRatio - 4f / 3f)) {
            PRESETS_16_9
        } else {
            PRESETS_4_3
        }
    }

    fun videoLayersFromEncodings(
        trackWidth: Int,
        trackHeight: Int,
        encodings: List<RtpParameters.Encoding>,
        isSVC: Boolean,
    ): List<LivekitModels.VideoLayer> {
        // Calculate aspect ratio to maintain it across resolutions
        val aspectRatio = trackWidth.toFloat() / trackHeight.toFloat()
        
        // Define fixed resolutions
        val resolutions = listOf(
            Pair(1920, 1080), // 1080p
            Pair(1280, 720),  // 720p
            Pair(960, 540)    // 540p
        )

        return resolutions.map { (width, height) ->
            LivekitModels.VideoLayer.newBuilder().apply {
                this.width = width
                this.height = height
                quality = when (height) {
                    1080 -> LivekitModels.VideoQuality.HIGH
                    720 -> LivekitModels.VideoQuality.MEDIUM
                    540 -> LivekitModels.VideoQuality.LOW
                    else -> LivekitModels.VideoQuality.HIGH
                }
                // Use a reasonable bitrate for each resolution
                bitrate = when (height) {
                    1080 -> 3000000  // 3 Mbps for 1440p
                    720 -> 2000000   // 2 Mbps for 720p
                    540 -> 1000000   // 1 Mbps for 540p
                    else -> 3000000
                }
                ssrc = 0
            }.build()
        }
    }

    fun videoQualityForRid(rid: String): LivekitModels.VideoQuality {
        return when (rid) {
            "f" -> LivekitModels.VideoQuality.HIGH
            "h" -> LivekitModels.VideoQuality.MEDIUM
            "q" -> LivekitModels.VideoQuality.LOW
            else -> LivekitModels.VideoQuality.UNRECOGNIZED
        }
    }

    fun ridForVideoQuality(quality: LivekitModels.VideoQuality): String? {
        return when (quality) {
            LivekitModels.VideoQuality.HIGH -> "f"
            LivekitModels.VideoQuality.MEDIUM -> "h"
            LivekitModels.VideoQuality.LOW -> "q"
            else -> null
        }
    }
}
