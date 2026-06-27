---
source: "https://arxiv.org/abs/2312.02152v2"
title: "Steerers: A framework for rotation equivariant keypoint descriptors"
author: "Georg Bökman, Johan Edstedt, Michael Felsberg, Fredrik Kahl"
year: "2023"
publication: "arXiv preprint / cs.CV"
download: "https://arxiv.org/pdf/2312.02152v2"
pdf: "https://arxiv.org/pdf/2312.02152v2"
captured_at: "2026-06-18T11:15:55Z"
updated_at: "2026-06-18T11:15:55Z"
capture_tool: "scrapem"
source_name: "arxiv"
keyword: "ルカーチ・ジェルジュ"
query: "Georg Lukács"
tags:
  - "現代思想"
  - "マルクス主義"
  - "物象化論"
  - "西洋マルクス主義"
status: raw
---

# Steerers: A framework for rotation equivariant keypoint descriptors

- 著者: Georg Bökman, Johan Edstedt, Michael Felsberg, Fredrik Kahl
- 年: 2023
- 掲載情報: arXiv preprint / cs.CV
- 情報源: [arxiv](https://arxiv.org/abs/2312.02152v2)
- ダウンロード: https://arxiv.org/pdf/2312.02152v2
- PDF: https://arxiv.org/pdf/2312.02152v2

## Obsidian Links

- 研究動向: [[研究動向/ルカーチ・ジェルジュ-現代研究動向|ルカーチ・ジェルジュ-現代研究動向]]
- キーワード: [[ルカーチ・ジェルジュ]]
- 関連分野: [[現代思想]]
- 関連分野: [[マルクス主義]]
- 関連分野: [[物象化論]]
- 関連分野: [[西洋マルクス主義]]
- 関連タグ: #現代思想 #マルクス主義 #物象化論 #西洋マルクス主義

## Abstract

Image keypoint descriptions that are discriminative and matchable over large changes in viewpoint are vital for 3D reconstruction. However, descriptions output by learned descriptors are typically not robust to camera rotation. While they can be made more robust by, e.g., data augmentation, this degrades performance on upright images. Another approach is test-time augmentation, which incurs a significant increase in runtime. Instead, we learn a linear transform in description space that encodes rotations of the input image. We call this linear transform a steerer since it allows us to transform the descriptions as if the image was rotated. From representation theory, we know all possible steerers for the rotation group. Steerers can be optimized (A) given a fixed descriptor, (B) jointly with a descriptor or (C) we can optimize a descriptor given a fixed steerer. We perform experiments in these three settings and obtain state-of-the-art results on the rotation invariant image matching benchmarks AIMS and Roto-360. We publish code and model weights at https://github.com/georg-bn/rotation-steerers.

## Notes

- 自動収集された未処理ノート。正式ノート化する前に内容と出典を確認する。
