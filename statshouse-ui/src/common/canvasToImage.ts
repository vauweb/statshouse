// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

export function linkToImage(link: string): Promise<HTMLImageElement> {
  return new Promise((resolve) => {
    const img = new Image();
    img.onload = () => {
      resolve(img);
    };
    img.src = link;
  });
}

export async function toBlob(canvas: HTMLCanvasElement | OffscreenCanvas): Promise<Blob | null> {
  if (typeof window.OffscreenCanvas === 'function' && canvas instanceof window.OffscreenCanvas) {
    return canvas.convertToBlob({ type: 'image/png' });
  }
  return new Promise((resolve) => {
    (canvas as HTMLCanvasElement).toBlob?.(resolve, 'image/png');
  });
}
export function createCanvas(width: number, height: number): HTMLCanvasElement | OffscreenCanvas {
  if (typeof window.OffscreenCanvas === 'function' && width > 0 && height > 0) {
    return new window.OffscreenCanvas(width, height);
  }
  const ext = document.createElement('canvas');
  ext.width = width;
  ext.height = height;
  return ext;
}

export async function canvasToImageData(
  canvas: HTMLCanvasElement,
  x: number,
  y: number,
  width: number,
  height: number,
  result_width?: number,
  result_height?: number
): Promise<string> {
  const ext_width = result_width ?? width;
  const ext_height = result_height ?? (result_width ? Math.round((result_width / width) * height) : height);
  const ext = createCanvas(ext_width, ext_height);
  const ctx = ext.getContext('2d') as CanvasRenderingContext2D | OffscreenCanvasRenderingContext2D | null;
  if (ctx) {
    ctx.drawImage(canvas, x, y, width, height, 0, 0, ext_width, ext_height);
  }
  const blobExt = await toBlob(ext);
  if (!blobExt) {
    return '';
  }
  return URL.createObjectURL(blobExt);
}

export async function setBackgroundColor(
  data: string,
  color: string,
  result_width?: number,
  result_height?: number
): Promise<string> {
  if (!data || !color) {
    return '';
  }
  const imgElement = await linkToImage(data);
  if (!imgElement) {
    return '';
  }
  const { width, height } = imgElement;
  const ext_width = result_width ?? width;
  const ext_height = result_height ?? (result_width ? Math.round((result_width / width) * height) : height);
  const ext = createCanvas(ext_width, ext_height);
  const ctx = ext.getContext('2d') as CanvasRenderingContext2D | OffscreenCanvasRenderingContext2D | null;
  if (ctx) {
    ctx.globalAlpha = 1;
    ctx.fillStyle = color;
    ctx.fillRect(0, 0, ext_width, ext_height);
    ctx.drawImage(imgElement, 0, 0, width, height, 0, 0, ext_width, ext_height);
  }
  const blobExt = await toBlob(ext);
  if (!blobExt) {
    return '';
  }
  return URL.createObjectURL(blobExt);
}
